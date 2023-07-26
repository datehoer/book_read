import axios from "axios"

const api = axios.create({
    baseURL: "http://localhost:8089",
    withCredentials: true,
    headers: {
        "Content-Type": "application/json",
    },
})

const ApiService = {
    register: (username, password) => api.post("/api/register/", { username, password }),
    login: (username, password) => api.post("/api/login/", { username, password }),
    getBooks: (page) => api.get("/api/books/", { params: { page } }),
    searchBooks: (search_keyword) => api.get("/api/search/", { params: { search_keyword } }),
    getBookDetails: (book_id) => api.get("/api/book/", { params: { book_id } }),
    getArticleDetails: (article_id) => api.get("/api/article/", { params: { article_id } }),
}

export default ApiService
