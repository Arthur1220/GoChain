<template>
  <div class="spinner"></div>
</template>

<style scoped>
.spinner {
  width: 24px;
  height: 24px;
  border: 3px solid rgba(255,255,255,0.1);
  border-top-color: var(--color-primary);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }
</style>