import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Link } from 'react-router-dom';
import ApiService from "../../api";
import "./index.css";

function ArticleDetail() {
    const { articleId } = useParams();
    const [article, setArticle] = useState(null);
    const [prevArticleId, setPrevArticleId] = useState(null);
    const [nextArticleId, setNextArticleId] = useState(null);

    useEffect(() => {
        const fetchArticleDetail = async () => {
            try {
                const response = await ApiService.getArticleDetails(articleId);
                setArticle(response.data);
                const articles = JSON.parse(localStorage.getItem('articles'));
                const index = articles.findIndex(article => article.id === response.data.id);
                // 设置上一章和下一章的 ID
                if (index > 0) {
                    setPrevArticleId(articles[index - 1].id);
                }
                if (index < articles.length - 1) {
                    setNextArticleId(articles[index + 1].id);
                }
                window.scrollTo(0, 0);
            } catch (error) {
                console.error("Failed to fetch article details: ", error);
            }
        };

        fetchArticleDetail();
    }, [articleId]);

    if (!article) {
        return <div>Loading...</div>;
    }

    return (
        <div className="article-detail">
            <h1>{article.title}</h1>
            <p>文章ID: {article.id}</p>
            <p>书籍ID: {article.book_id}</p>
            <div dangerouslySetInnerHTML={{ __html: article.content }}></div>
            <p>编号: {article.num}</p>
            <div className="nav-buttons">
                {prevArticleId && <Link className="prev" to={`/article/${prevArticleId}`}>上一章</Link>}
                {nextArticleId && <Link className="next" to={`/article/${nextArticleId}`}>下一章</Link>}
            </div>
        </div>
    );
}

export default ArticleDetail;