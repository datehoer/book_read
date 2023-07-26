import React, { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import ApiService from "../../api"
import { Link } from 'react-router-dom';
import "./index.css"

function BookDetail() {
    const { bookId } = useParams()
    const [book, setBook] = useState(null)
    const [articles, setArticles] = useState([])

    useEffect(() => {
        const fetchBookDetail = async () => {
            try {
                const response = await ApiService.getBookDetails(bookId)
                setBook(response.data.book)
                setArticles(response.data.article)
                localStorage.setItem('articles', JSON.stringify(response.data.article))
            } catch (error) {
                console.error("Failed to fetch book details: ", error)
            }
        }

        fetchBookDetail()
    }, [bookId])

    if (!book || !articles) {
        return <div>Loading...</div>
    }

    return (
        <div className="book-detail">
            <h1>{book.book_name}</h1>
            <p>书籍ID: {book.book_id}</p>
            <p>章节数: {book.article_count}</p>
            <p>标签: {book.tag}</p>

            <h2>章节列表：</h2>
            <ul>
                {articles.map((article, index) => (
                    <li key={index}>
                        <Link to={`/article/${article.id}`}>
                            <h3>{article.title}</h3>
                        </Link>
                        <p>章节ID: {article.id}</p>
                        <p>书籍ID: {article.book_id}</p>
                        <p>编号: {article.num}</p>
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default BookDetail
