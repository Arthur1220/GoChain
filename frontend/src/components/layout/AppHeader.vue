<script setup>
import { RouterLink, useRoute } from 'vue-router';
import { Terminal, LayoutDashboard, ShieldAlert } from 'lucide-vue-next';

const route = useRoute();
const isActive = (path) => route.path === path;
</script>

<template>
  <header class="app-header">
    <div class="container">
      <div class="brand">
        <div class="logo-icon"><Terminal :size="20" /></div>
        <span class="logo-text">Go-Chain_Turbo</span>
      </div>

      <nav class="nav">
        <RouterLink to="/" class="nav-link" :class="{ active: isActive('/') }">
          Home
        </RouterLink>
        <RouterLink to="/dashboard" class="nav-link" :class="{ active: isActive('/dashboard') }">
          Dashboard
        </RouterLink>
        <RouterLink to="/admin" class="nav-link" :class="{ active: isActive('/admin') }">
          Admin
        </RouterLink>
      </nav>
    </div>
  </header>
</template>

<style scoped>
.app-header {
  border-bottom: 1px solid var(--border-color);
  background: rgba(9, 9, 11, 0.8);
  backdrop-filter: blur(8px);
  position: sticky;
  top: 0;
  z-index: 10;
}

.container {
  max-width: var(--container-width);
  margin: 0 auto;
  padding: 0 2rem;
  height: 70px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.brand { display: flex; align-items: center; gap: 10px; color: var(--color-primary); }
.logo-icon { background: rgba(34, 197, 94, 0.1); padding: 6px; border-radius: 6px; display: flex; }
.logo-text { font-family: var(--font-mono); font-weight: 700; letter-spacing: -0.5px; color: var(--text-main); }

.nav { display: flex; gap: 2rem; }
.nav-link { font-size: 0.9rem; font-weight: 500; color: var(--text-muted); transition: 0.2s; }
.nav-link:hover { color: var(--text-main); }
.nav-link.active { color: var(--color-primary); font-weight: 700; }
</style>