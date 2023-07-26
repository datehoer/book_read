import { BrowserRouter as Router, Route, Routes } from "react-router-dom"
import React from "react"
import PrivateRoute from './components/PrivateRoute' // 引入PrivateRoute
import Login from "./components/Login"
import BookList from "./components/BookList"
import BookDetail from "./components/BookDetail"
import ArticleDetail from "./components/ArticleDetail"

function App() {
    return (
        <Router>
            <Routes>
                <Route path='/login' element={<Login />} />
                <Route path='/' element={<PrivateRoute><BookList /></PrivateRoute>} />
                <Route path='/book/:bookId' element={<PrivateRoute><BookDetail /></PrivateRoute>} />
                <Route path='/article/:articleId' element={<PrivateRoute><ArticleDetail /></PrivateRoute>} />
            </Routes>
        </Router>
    )
}

export default App
