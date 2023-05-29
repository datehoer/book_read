<template>
  <div>
    <h1>图书列表</h1>
    <ul>
      <li v-for="book in books" :key="book.book_id">
        <router-link :to="{ name: 'book', params: { id: book.book_id } }">
          {{ book.book_name }}
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from "axios";

export default {
  name: "HomeView",
  //   components: {
  //     HelloWorld
  //   }
  data() {
    return {
      books: [],
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
          this.books = response.data;
          console.log(response.data)
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>
