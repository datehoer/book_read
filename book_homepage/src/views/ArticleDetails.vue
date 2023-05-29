<template>
    <div>
      <h1>{{ article.title }}</h1>
      <p>作者：{{ article.author }}</p>
      <p v-html="article.content"></p>
      <!-- 其他文章详细信息 -->
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        article: {}
      };
    },
    created() {
      this.getArticleDetails();
    },
    methods: {
      getArticleDetails() {
        const articleId = this.$route.params.id;
        axios.defaults.baseURL='/api'
        axios.get(`/article?article_id=${articleId}`)
          .then(response => {
            this.article = response.data;
          })
          .catch(error => {
            console.error(error);
          });
      }
    }
  };
  </script>
  