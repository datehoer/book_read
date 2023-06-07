<template>
  <div>
    <input type="text" name="" id="" placeholder="关键词" v-model="keyword">
    <button type="primary" @click="search">搜索</button>
    <h1>图书列表</h1>
    <ul>
      <li v-for="book in books" :key="book.book_id">
        <router-link :to="{ name: 'book', params: { id: book.book_id } }">
          {{ book.book_name }}
        </router-link>
      </li>
    </ul>
    <!-- 上一页按钮 -->
    <button v-if="currentPage > 1" @click="previousPage">上一页</button>

    <!-- 下一页按钮 -->
    <button v-if="currentPage < pageCount" @click="nextPage">下一页</button>

    <!-- 显示当前页码和总页数 -->
    <p>当前页码: {{ currentPage }}</p>
    <p>总页数: {{ pageCount }}</p>

  </div>
</template>

<script>
// @ is an alias to /src
import axios from "axios";

export default {
  name: "HomeView",
  data() {
    return {
      books: [],
      currentPage: 1,
      pageCount: 0,
      keyword: "",
    };
  },
  created() {
    this.getAllBooks();
  },
  methods: {
    getAllBooks() {
      axios.defaults.baseURL='/api'
      axios.get("/books")
        .then((response) => {
          this.books = response.data.books;
          this.pageCount = response.data.pageCount;
          console.log(this.books)
        })
        .catch((error) => {
          console.error(error);
        });
    },
    getMoreBooks(){
        axios.defaults.baseURL='/api'
        axios.get("/books?page="+this.currentPage)
            .then((response) => {
            this.books = response.data.books;
            this.pageCount = response.data.pageCount;
            console.log(this.books)
            })
            .catch((error) => {
            console.error(error);
            });
    },
    previousPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.getMoreBooks();
      }
    },
    nextPage() {
      if (this.currentPage < this.pageCount) {
        this.currentPage++;
        this.getMoreBooks();
      }
    },
    search(){
        axios.defaults.baseURL='/api'
        axios.get("/search?search_keyword="+this.keyword)
            .then((response) => {
            this.books = response.data;
            this.pageCount = 0;
            console.log(this.books)
            })
            .catch((error) => {
            console.error(error);
            });
    }
  },
};
</script>
