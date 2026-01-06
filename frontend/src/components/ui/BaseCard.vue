<template>
  <div class="base-card">
    <slot></slot>
  </div>
</template>

<style scoped>
.base-card {
  background-color: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: 1.5rem;
}
</style>