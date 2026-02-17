import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../pages/Dashboard.vue'
import Members from '../pages/Members.vue'
import Plans from '../pages/Plans.vue'
import Classes from '../pages/Classes.vue'
import Billing from '../pages/Billing.vue'
import MemberProfile from '../pages/MemberProfile.vue'

const routes = [
  { path: '/', component: Dashboard },
  { path: '/members', component: Members },
  { path: '/members/:id', component: MemberProfile, props: true },
  { path: '/plans', component: Plans },
  { path: '/classes', component: Classes },
  { path: '/billing', component: Billing },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
