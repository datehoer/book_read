<template>
    <div>
    <van-search v-model="value" placeholder="请输入搜索关键词" />
      <ul>   
        <li v-for="book in books" :key="book.book_id">
          <router-link :to="{ name: 'book', params: { id: book.book_id } }">
            {{ book.book_name }}
          </router-link>
        </li>
      </ul>
      <van-pagination v-model="currentPage" :page-count="pageCount" mode="simple" @click="getMoreBooks"/>
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
            console.log(response.data)
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
              console.log(response.data)
              })
              .catch((error) => {
              console.error(error);
              });
      }
    },
  };
  </script>
  