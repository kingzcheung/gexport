import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: "Home",
            component: () => import("../views/Home.vue")
        },
        {
            path: '/sql',
            name: "Sql",
            props: true,
            component: () => import("../views/Sql.vue")
        },

    ],
})


export default router;