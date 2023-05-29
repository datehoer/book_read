// module.exports = {
//     server: {
//       proxy: {
//         '/login': {
//           target: 'http://127.0.0.1:8080',
//           changeOrigin: true,
//           rewrite: (path) => path.replace(/^\/login/, '')
//         },
//         '/register': {
//           target: 'http://127.0.0.1:8080',
//           changeOrigin: true,
//           rewrite: (path) => path.replace(/^\/register/, '')
//         }
//       }
//     }
//   }  
export default defineConfig({
    plugins: [vue()],
    server: {
        host: '0.0.0.0'
    }
})