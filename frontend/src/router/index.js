import { createRouter, createWebHistory } from 'vue-router'

// Importamos o Dashboard diretamente pois ele já existe
import DashboardView from '../views/DashboardView.vue'

// Importamos as novas Views dinamicamente (Lazy Loading)
// Isso é bom para performance e evita erros se o arquivo ainda não existir fisicamente
const HomeView = () => import('../views/HomeView.vue')
const AdminView = () => import('../views/AdminView.vue')

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView
    },
    {
      // Rota "Escondida" de Admin
      path: '/admin',
      name: 'admin',
      component: AdminView
    }
  ]
})

export default router