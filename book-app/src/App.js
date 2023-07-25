import { BrowserRouter as Router, Route, Routes  } from 'react-router-dom';
import React from 'react';
import Login from './components/Login';
import BookList from './components/BookList';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/" element={<BookList />} exact />
      </Routes>
    </Router>
  );
}

export default App;