import { createRouter, createWebHistory } from "vue-router"
import HomeView from "../views/HomeView.vue"
import BookDetails from "../views/BookDetails.vue"
import ArticleDetails from "../views/ArticleDetails.vue"
const routes = [
    {
        path: "/",
        name: "home",
        component: HomeView,
    },
    {
        path: "/login",
        name: "login",
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ "../views/LoginView.vue"),
    },
    {
        path: "/about",
        name: "about",
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ "../views/AboutView.vue"),
    },
    {
        path: "/book/:id",
        name: "book",
        component: BookDetails,
    },
    {
        path: "/article/:id",
        name: "article",
        component: ArticleDetails,
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
})

router.beforeEach((to, from, next) => {
    const isLoggedIn = checkIfUserIsLoggedIn("auth") // 你可以根据实际情况检查用户是否已登录
    if (to.name !== "login" && !isLoggedIn) {
        next({
            name: "login",
            query: {
                redirect: to.path
            }
        }) // 如果用户未登录且访问的页面不是登录页，则跳转到登录页
    } else {
        next() // 继续导航到目标页面
    }
})

function checkIfUserIsLoggedIn(cookieName) {
    const cookies = document.cookie.split(";")
    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim()
        if (cookie.startsWith(cookieName + "=")) {
            return true
        }
    }
    return false
}

export default router
