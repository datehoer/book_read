<template>
  <van-form @submit="login">
    <van-cell-group inset>
      <van-field v-model="username" name="用户名" label="用户名" placeholder="用户名"
        :rules="[{ required: true, message: '请填写用户名' }]" />
      <van-field v-model="password" type="password" name="密码" label="密码" placeholder="密码"
        :rules="[{ required: true, message: '请填写密码' }]" />
    </van-cell-group>
    <div style="margin: 16px">
      <van-button round block type="primary" native-type="submit">
        提交
      </van-button>
    </div>
  </van-form>
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
  