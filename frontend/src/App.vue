<script setup>
import { RouterView } from 'vue-router';
import AppHeader from './components/layout/AppHeader.vue';
import AppFooter from './components/layout/AppFooter.vue';
</script>

<template>
  <div class="layout-wrapper">
    <AppHeader />
    
    <main class="main-content">
      <div class="container">
        <RouterView v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </RouterView>
      </div>
    </main>

    <AppFooter />
  </div>
</template>

<style scoped>
.layout-wrapper { display: flex; flex-direction: column; min-height: 100vh; }
.main-content { flex: 1; padding: 3rem 0; }
.container { max-width: var(--container-width); margin: 0 auto; padding: 0 2rem; width: 100%; }

/* Animação suave entre páginas */
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>