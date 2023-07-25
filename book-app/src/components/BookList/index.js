import React, { useEffect, useState } from 'react';
import ApiService from "../../api";
import "./index.css"

function BookList() {
    const [books, setBooks] = useState([]);
    const [page, setPage] = useState(1);
    const [pageCount, setPageCount] = useState(0);

    useEffect(() => {
        const fetchBooks = async () => {
            try {
                const response = await ApiService.getBooks(page);
                setBooks(response.data.books);
                setPageCount(response.data.pageCount);
            } catch (error) {
                console.error('Failed to fetch books: ', error);
            }
        };

        fetchBooks();
    }, [page]);

    const handleNextPage = () => {
        if (page < pageCount) {
            setPage(page + 1);
        }
    }

    const handlePrevPage = () => {
        if (page > 1) {
            setPage(page - 1);
        }
    }

    return (
        <div className="book-list">
            <h1>Books</h1>
            {books.map((book, index) => (
                <li key={index}>
                    <h2>{book.book_name}</h2>
                    <p>书籍ID: {book.book_id}</p>
                    <p>章节数: {book.article_count}</p>
                    <p>标签: {book.tag}</p>
                </li>
            ))}
            <div className="pagination-buttons">
                {page > 1 && <button onClick={handlePrevPage}>Previous Page</button>}
                {page < pageCount && <button onClick={handleNextPage}>Next Page</button>}
            </div>
        </div>
    );
}

export default BookList;