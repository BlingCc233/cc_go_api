import { createRouter, createWebHistory } from 'vue-router';
import LoginPage from '../views/Login.vue';
import MainPage from '../views/Main.vue';
import HelloWorld from '../components/HelloWorld.vue'
import Cookies from 'js-cookie';

const routes = [
  {
    path: '/',
    redirect: '/main'
  },
  {
    path: '/login',
    component: LoginPage,
    meta: { isLogin: true }
  },
  {
    path: '/main',
    component: MainPage,
    children: [
      {
        path: '/dbs',
        component: HelloWorld
      }
    ],
    meta: { requiresAuth: true } // 添加 requiresAuth 元信息
  },

];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !isAuthenticated()) {
    next('/login');
  }
  if (to.meta.isLogin && isAuthenticated()){
    next('/main');
  }
  else {
    next();
  }
});


// 模拟简单的身份验证逻辑，实际项目中需要根据具体情况实现
function isAuthenticated() {
  return Cookies.get('token') !== undefined;
}

export default router;
