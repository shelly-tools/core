import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'root',
    component: () => import('@/views/Dashboard.vue'),
    children: [
      {
        path: 'buildings',
        name: 'buildings',
        component: () => import('../views/Buildings.vue'),
      },
    ],
    /*
    beforeEnter: (to, from, next) => {
      const user = JSON.parse(localStorage.getItem('user'));
      if (user !== null && user !== undefined) {
        next('/');
      } else {
        next('/login');
      }
    },
    */
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/Login.vue'),
    beforeEnter: (to, from, next) => {
      const user = localStorage.getItem('user');
      if (user === null) {
        next();
      } else {
        next('/');
      }
    },
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
