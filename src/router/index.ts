import AuthView from '@/views/auth/AuthView.vue'
import CustomerView from '@/views/customer/CustomerView.vue'
import HomeView from '@/views/home/HomeView.vue'
import OrderView from '@/views/order/OrderView.vue'
import ProductView from '@/views/product/ProductView.vue'
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  { path: '/', component: () => HomeView },
  { path: '/auth', beforeEnter: () => { }, component: () => AuthView },
  {
    path: '/customer/:customer_number', children: [
      {
        path: '',
        name: "customer-details",
        component: CustomerView
      },
      {
        path: 'order/:order_number', // matches /customer/123/orders/999
        name: 'order-details',
        component: OrderView
      }
    ]
  },
  { path: '/product/:product_id', component: ProductView, name: "product-page" }
]
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
