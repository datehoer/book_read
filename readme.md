#### 接口文档
- 获取所有图书信息
- 请求方式：GET
- URL：/books
- 描述：用于获取所有图书的信息
#### 分页获取图书信息
- 请求方式：GET
- URL：/books?page={page_number}
- 描述：用于分页获取图书的信息，每页返回10条数据
- 参数：
- - page_number (必需)：页码，表示需要获取的页数
#### 获取书籍详情
- 请求方式：GET
- URL：/book?book_id={book_id}
- 描述：用于获取特定书籍的详细信息
- 参数：
- - book_id (必需)：书籍ID，表示需要获取的书籍的唯一标识
#### 获取文章详情
- 请求方式：GET
- URL：/article?article_id={article_id}
- 描述：用于获取特定文章的详细信息
- 参数：
- - article_id (必需)：文章ID，表示需要获取的文章的唯一标识