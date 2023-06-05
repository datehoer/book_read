const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})
module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://124.221.222.201:8089',
        // 修改为以下配置
        changOrigin: true,
        pathRewrite: {
          '^/api': '', // 去掉请求路径中的/api前缀
        },
      },
      '/book': {
        target: 'http://124.221.222.201:8089',
        // 修改为以下配置
        changOrigin: true,
        pathRewrite: {
          '^/book': '/book', // 去掉请求路径中的/api前缀
        },
      },
      '/article': {
        target: 'http://124.221.222.201:8089',
        // 修改为以下配置
        changOrigin: true,
        pathRewrite: {
          '^/article': '/article', // 去掉请求路径中的/api前缀
        },
      },
    },
  },
};
