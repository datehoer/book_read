package main

import (
	"book_api/db"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type booksHandler struct{}

type Book struct {
	BookName     string `json:"book_name"`
	BookId       string `json:"book_id"`
	ArticleCount string `json:"article_count"`
	Tag          string `json:"tag"`
}
type Article struct {
	ID      int    `json:"id"`
	BookID  string `json:"book_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Num     int    `json:"num"`
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查是否存在登录的认证 Cookie
		cookie, err := r.Cookie("auth")
		if err != nil || cookie.Value == "" {
			// 未登录，返回未授权的错误响应
			http.Error(w, "未登录", http.StatusUnauthorized)
			return
		}

		// 在这里进行其他权限验证逻辑，比如验证 Cookie 的合法性、权限等
		dbConn, err := db.ConnectDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dbConn.Close()

		res, err := dbConn.Query("select username from book_user where cookie = ?", cookie.Value)
		if err != nil {
			http.Error(w, "登录信息过期", http.StatusInternalServerError)
			return
		}
		defer res.Close()
		if res.Next() {
			var username string
			err := res.Scan(&username)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "登录信息过期", http.StatusInternalServerError)
			return
		}
	})
}

type BooksResponse struct {
	Books     []*Book `json:"books"`
	PageCount int     `json:"pageCount"`
}

func (h *booksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	offset := (page - 1) * 10
	query := "SELECT book_id, book_name, tag_name, COUNT(book_id) AS article_count FROM book GROUP BY book_id, book_name, tag_name LIMIT 10 OFFSET ? "
	selectAll := "SELECT count(DISTINCT book_id) from book"
	rows, err := dbConn.Query(query, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	booksResponse := BooksResponse{
		Books:     make([]*Book, 0),
		PageCount: 0,
	}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.BookId, &book.BookName, &book.Tag, &book.ArticleCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		booksResponse.Books = append(booksResponse.Books, &book)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count, err := dbConn.Query(selectAll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for count.Next() {
		var bookCount int
		err := count.Scan(&bookCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		booksResponse.PageCount = int(math.Ceil(float64(bookCount) / 10))
	}
	defer rows.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booksResponse)
}

type bookHandler struct{}

func (h *bookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	// 解析 URL 中的参数
	params := r.URL.Query()

	// 获取 book_id 参数
	bookID := params.Get("book_id")
	// 查询书籍的详细信息
	var book Book
	bookSql := "SELECT book_id, book_name, tag_name, COUNT(book_id) AS article_count FROM book WHERE book_id = ? GROUP BY book_id, book_name, tag_name"
	err = dbConn.QueryRow(bookSql, bookID).Scan(&book.BookId, &book.BookName, &book.Tag, &book.ArticleCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 查询书籍的章节信息
	articles := make([]Article, 0)
	rows, err := dbConn.Query("SELECT id, book_id, title, indexes FROM book WHERE book_id = ?", bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.BookID, &article.Title, &article.Num)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookData := struct {
		Book    Book      `json:"book"`
		Article []Article `json:"article"`
	}{
		Book:    book,
		Article: articles,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookData)
}

type articleHandler struct{}

func (h *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	// 解析 URL 中的参数
	params := r.URL.Query()

	// 获取 id 参数
	idStr := params.Get("article_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid chapter ID", http.StatusBadRequest)
		return
	}

	// 查询章节内容
	var article Article
	err = dbConn.QueryRow("SELECT id, book_id, title, change_content_html FROM book WHERE id = ?", id).Scan(&article.ID, &article.BookID, &article.Title, &article.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerHandler struct{}

func (h *registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "无效的请求方法", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()
	res, err := dbConn.Exec("insert into book_user(username, password) values (?, ?)", user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(rowsAffected)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("注册成功"))
}

type loginHandler struct{}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "无效的请求方法", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()
	res, err := dbConn.Query("select id, password from book_user where username = ?", user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Close()
	if res.Next() {
		var userID string
		var storedPassword string
		err := res.Scan(&userID, &storedPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if user.Password != storedPassword {
			http.Error(w, "用户名或密码不正确", http.StatusUnauthorized)
			return
		}
		authValue := user.Username + user.Password
		cookieValue := md5Hash(authValue)
		res, err := dbConn.Exec("update book_user set cookie = ? where id = ?", cookieValue, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(rowsAffected)
		cookie := &http.Cookie{
			Name:    "auth",
			Value:   cookieValue,
			Expires: time.Now().Add(24 * time.Hour), // 设置Cookie过期时间
			Path:    "/",
		}
		http.SetCookie(w, cookie)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("登录成功"))
	} else {
		http.Error(w, "用户名或密码不正确", http.StatusUnauthorized)
	}
}

func md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置允许的源
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		// 设置允许的请求方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// 设置允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// 允许携带 cookie
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			// 如果是 OPTIONS 请求，直接返回成功响应
			w.WriteHeader(http.StatusOK)
			return
		}

		// 继续处理其他请求
		next.ServeHTTP(w, r)
	})
}

type searchHandler struct{}

func (h *searchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	// 解析 URL 中的参数
	params := r.URL.Query()

	// 获取 book_id 参数
	keyWord := params.Get("search_keyword")
	// 查询书籍的详细信息
	bookSql := "SELECT book_id, book_name, tag_name, COUNT(book_id) AS article_count FROM book WHERE book_name like ? GROUP BY book_id, book_name, tag_name"
	rows, err := dbConn.Query(bookSql, "%"+keyWord+"%")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	booksResponse := BooksResponse{
		Books:     make([]*Book, 0),
		PageCount: 0,
	}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.BookId, &book.BookName, &book.Tag, &book.ArticleCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		booksResponse.Books = append(booksResponse.Books, &book)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booksResponse)
}

func main() {
	booksHandler := &booksHandler{}
	bookHandler := &bookHandler{}
	articleHandler := &articleHandler{}
	registerHandler := &registerHandler{}
	loginHandler := &loginHandler{}
	searchHandler := &searchHandler{}
	http.Handle("/api/register", registerHandler)
	http.Handle("/api/login", corsMiddleware(loginHandler))
	http.Handle("/api/books", authMiddleware(booksHandler))
	http.Handle("/api/search/", http.StripPrefix("/api/search", authMiddleware(searchHandler)))
	http.Handle("/api/book/", http.StripPrefix("/api/book", authMiddleware(bookHandler)))
	http.Handle("/api/article/", http.StripPrefix("/api/article", authMiddleware(articleHandler)))
	log.Fatal(http.ListenAndServe(":8089", nil))
}
