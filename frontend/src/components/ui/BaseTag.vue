<template>
  <span class="base-tag">
    <slot></slot>
  </span>
</template>

<style scoped>
.base-tag {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--border-color);
  padding: 0.2rem 0.6rem;
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: 700;
  color: var(--text-muted);
  font-family: monospace;
}
</style>