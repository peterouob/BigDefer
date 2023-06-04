import {createWebHistory,createRouter} from "vue-router"
import Login from '../view/Login.vue'

const routes = [
    {
        path:'/login',
        name:'login',
        component:Login
    }
]

const router = createRouter({
    history:createWebHistory(),
    routes
})

export default router