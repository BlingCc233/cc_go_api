import {createRouter, createWebHistory} from 'vue-router';
import LoginPage from '../views/Login.vue';
import MainPage from '../views/Main.vue';
import HelloWorld from '../components/HelloWorld.vue'
import Home from '../components/Home.vue'
import Cookies from 'js-cookie';
import swiper from '../pages/swiper.vue';
import NotFount from '../pages/404.vue'
import WishList from '../components/WishList.vue'
import MainApi from "../pages/MainApi.vue";

const routes = [
    {
        path: '/',
        redirect: '/home',
        meta: {
            title: '首页说是',
        },
    },
    {
        path: '/404',
        component: NotFount,
        meta: {
            title: '404',
        },
    },
    {
        path: '/login',
        component: LoginPage,
        meta: {isLogin: true, title: '登录'}
    },
    {
        path: '/main',
        component: MainPage,
        children: [
            {
                path: '',
                component: Home,
                meta: {
                    title: '首页说是',
                },
            },
            {
                path: '/dbs',
                component: HelloWorld,
                meta: {
                    title: 'dbs',
                },
            },
            {
                path: '/home',
                component: Home,
                meta: {
                    title: '首页说是',
                },
            },
            {
                path: '/test',
                component: swiper,
                meta: {
                    title: 'test',
                },
            },
            {
                path: '/:pathMatch(.*)*',
                redirect: '/404',
                meta: {
                    title: '404',
                },
            },
            {
                path: '/wishlist',
                component: WishList,
                meta: {
                    title: 'wishlist',
                },
            }

        ],
        meta: {requiresAuth: true} // 添加 requiresAuth 元信息
    },
    {
        path: '/api',
        component: MainApi,
        meta: {
            title: 'api'
        },
        children: [
            {
                path: '',
                component: Home,
                meta: {
                    title: 'api',
                },
            },
            {
                path: '/api/home',
                component: Home,
                meta: {
                    title: 'api',
                },
            },
            {
                path: '/api/test',
                component: swiper,
                meta: {
                    title: 'test',
                },
            },
            {
                path: '/:pathMatch(.*)*',
                redirect: '/404',
                meta: {
                    title: '404',
                },
            },
            {
                path: '/api/dbs',
                component: HelloWorld,
                meta: {
                    title: 'dbs',
                },
            },
        ]

    }

];

const router = createRouter({
    history: createWebHistory(),
    routes
});

router.beforeEach((to, from, next) => {

    if (to.meta.requiresAuth && !isAuthenticated()) {
        next('/login');
    }
    if (to.meta.isLogin && isAuthenticated()) {
        next('/home');
    }
    else {
        document.title = to.meta.title;
        next();
    }
});


// 模拟简单的身份验证逻辑，实际项目中需要根据具体情况实现
function isAuthenticated() {
    if (Cookies.get('token') === undefined) {
        return false;
    }
    const token = Cookies.get('token');
    const user = Cookies.get('user');
    let decodedToken = {
        username: '-1'
    };
    // 解码Token，
    try {
        decodedToken = JSON.parse(atob(token.split('.')[1]));
    } catch (e) {
        return false;
    }
    return user === decodedToken.username;

}

export default router;
