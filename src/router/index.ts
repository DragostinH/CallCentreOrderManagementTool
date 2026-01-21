import AuthView from '@/views/auth/AuthView.vue'
import CustomerView from '@/views/customer/CustomerView.vue'
import HomeView from '@/views/home/HomeView.vue'
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  { path: '/', component: () => HomeView },
  { path: '/auth', beforeEnter: () => {}, component: () => AuthView },
  { path: '/customer/:customerId', component: () => CustomerView },
]
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
