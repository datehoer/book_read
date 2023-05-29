<template>
  <div>
    <h1>图书列表</h1>
    <ul>
      <li v-for="book in books" :key="book.id">
        <router-link :to="{ name: 'BookDetails', params: { id: book.id } }">
          {{ book.title }}
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
      axios
        .get("http://127.0.0.1:8080/books")
        .then((response) => {
          this.books = response.data;
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>
