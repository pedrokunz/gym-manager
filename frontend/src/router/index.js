import { createRouter, createWebHistory } from 'vue-router'
import Members from '../pages/Members.vue'
import Plans from '../pages/Plans.vue'
import Classes from '../pages/Classes.vue'
import Billing from '../pages/Billing.vue'

const routes = [
  { path: '/', component: Members },
  { path: '/plans', component: Plans },
  { path: '/classes', component: Classes },
  { path: '/billing', component: Billing },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
