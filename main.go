package main

import (
	"book_api/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func (h *booksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	offset := (page - 1) * 10
	query := "SELECT book_id, book_name, tag_name, COUNT(book_id) AS article_count FROM book GROUP BY book_id, book_name, tag_name LIMIT 10 OFFSET ? "
	rows, err := dbConn.Query(query, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	books := make([]*Book, 0)
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.BookId, &book.BookName, &book.Tag, &book.ArticleCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
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

	// 获取 book_id 参数
	bookID := params.Get("book_id")

	// 获取 id 参数
	idStr := params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid chapter ID", http.StatusBadRequest)
		return
	}

	// 查询章节内容
	var article Article
	err = dbConn.QueryRow("SELECT id, book_id, title, content_html FROM book WHERE book_id = ? and  id = ?", bookID, id).Scan(&article.ID, &article.BookID, &article.Title, &article.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func main() {
	http.Handle("/books", &booksHandler{})
	http.Handle("/book/", http.StripPrefix("/book", &bookHandler{}))
	http.Handle("/article/", http.StripPrefix("/article", &articleHandler{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
