import React, { useEffect, useState, useCallback  } from "react"
import ApiService from "../../api"
import { Link } from "react-router-dom"
import "./index.css"

function BookList() {
    const [books, setBooks] = useState([])
    const [page, setPage] = useState(1)
    const [pageCount, setPageCount] = useState(0)
    const [searchKeyword, setSearchKeyword] = useState("")

    const fetchBooks = useCallback(async () => {
        try {
            let response;
            if (searchKeyword) {
                response = await ApiService.searchBooks(searchKeyword);
            } else {
                response = await ApiService.getBooks(page);
            }
            setBooks(response.data.books);
            setPageCount(response.data.pageCount);
            window.scrollTo(0, 0);
        } catch (error) {
            console.error('Failed to fetch books: ', error);
        }
    }, [page, searchKeyword]);

    useEffect(() => {
        fetchBooks();
    }, [fetchBooks]);

    const handleSearchBooks = (event) => {
        event.preventDefault();
        fetchBooks();
    }

    const handleNextPage = () => {
        if (page < pageCount) {
            setPage(page + 1)
        }
    }

    const handlePrevPage = () => {
        if (page > 1) {
            setPage(page - 1)
        }
    }

    return (
        <div className='book-list'>
            <h1>Books</h1>
            <form onSubmit={handleSearchBooks}>
                <input
                    type='text'
                    value={searchKeyword}
                    onChange={(e) => setSearchKeyword(e.target.value)}
                    placeholder='Search books...'
                />
                {/* <button type='submit'>Search</button> */}
            </form>
            {books.map((book, index) => (
                <li key={index}>
                    <Link to={`/book/${book.book_id}`}>
                        <h2>{book.book_name}</h2>
                    </Link>
                    <p>书籍ID: {book.book_id}</p>
                    <p>章节数: {book.article_count}</p>
                    <p>标签: {book.tag}</p>
                </li>
            ))}
            <div className='pagination-buttons'>
                {page > 1 && <button onClick={handlePrevPage}>Previous Page</button>}
                {page < pageCount && <button onClick={handleNextPage}>Next Page</button>}
            </div>
        </div>
    )
}

export default BookList
