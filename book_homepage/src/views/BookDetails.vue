<template>
  <div>
    <button type="primary" @click="back">返回上一页</button>
    <h1>{{ book.book.book_name }}</h1>
    <p v-if="book.book.tag">标签：{{ book.book.tag }}</p>
    <p v-if="book.book.article_count">
      章节数：{{ book.book.article_count }}章
    </p>
    <ul>
      <li v-for="article in book.article" :key="article.id">
        <router-link
          v-if="article.id"
          :to="{ name: 'article', params: { id: article.id } }"
        >
          {{ article.title }}
        </router-link>
      </li>
    </ul>
    <!-- 其他书籍详细信息 -->
  </div>
</template>
  
<script>
import axios from "axios";

export default {
  name: "BookDetails",
  data() {
    return {
      book: {
        book: {
          book_name: "",
          tag: "",
          article_count: "",
        },
        article: [
          {
            id: "",
            title: "",
          },
        ],
      },
    };
  },
  created() {
    this.getBookDetails();
  },
  methods: {
    getBookDetails() {
      const bookId = this.$route.params.id;
      axios.defaults.baseURL = "/api";
      let get_url = "/book?book_id=" + bookId;
      axios
        .get(get_url)
        .then((response) => {
          this.book = response.data;
          console.log(response.data);
        })
        .catch((error) => {
          console.error(error);
        });
    },
    back() {
      if (window.history.length <= 1) {
        this.$router.push({ path: "/" });
        return false;
      } else {
        this.$router.go(-1);
      }
    },
  },
};
</script>
