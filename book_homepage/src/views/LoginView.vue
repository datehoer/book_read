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
  import axios from "axios";
  export default {
    data() {
      return {
        username: '',
        password: '',
      };
    },
    methods: {
      login() {
        const credentials = {
        username: this.username,
        password: this.password,
      };
      axios.defaults.baseURL='/api'
      // 发送登录请求
      axios.create({
        withCredentials: true,
      }).post('/login', credentials)
        .then(response => {
          // 处理登录成功的情况
          // 根据后端返回的响应进行相应的处理
          console.log(response);
          // 其他操作，例如页面跳转等
        })
        .catch(error => {
          // 处理登录失败的情况
          // 根据后端返回的错误信息进行相应的处理
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
  