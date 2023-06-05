<template>
  <div class="login">
    <h2>Login</h2>
    <form @submit.prevent="login">
      <div>
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="username" required>
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required>
      </div>
      <button type="submit">Login</button>
    </form>
  </div>
</template>
  
<script>
// import { useRouter } from 'vue-router';
import axios from "axios";
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    login() {
      const credentials = {
        username: this.username,
        password: this.password,
      };
      axios.defaults.baseURL = "/api";
      const redirect = this.$route.query.redirect;
      // 发送登录请求
      axios
        .create({
          withCredentials: true,
        })
        .post("/login", credentials)
        .then((response) => {
          console.log(response);
          if (redirect) {
            this.$router.push(redirect);
          } else {

            this.$router.push({
              name: "home",
            });
          }
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>
  
<style scoped>
.login {
  max-width: 300px;
  margin: 0 auto;
}
</style>
  