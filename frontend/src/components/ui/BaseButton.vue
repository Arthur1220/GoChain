<script setup>
defineProps({
  variant: { type: String, default: 'primary' }, // primary, danger, outline
  disabled: Boolean,
  loading: Boolean
});
</script>

<template>
  <button :class="['base-btn', `btn-${variant}`]" :disabled="disabled || loading">
    <span v-if="loading">...</span>
    <slot v-else></slot>
  </button>
</template>

<style scoped>
.base-btn {
  padding: 0.75rem 1.5rem;
  border-radius: var(--radius-md);
  font-weight: 600;
  cursor: pointer;
  border: 1px solid transparent;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  transition: 0.2s;
}
.base-btn:disabled { opacity: 0.6; cursor: not-allowed; }

.btn-primary { background: var(--color-primary); color: var(--text-inverse); }
.btn-primary:hover:not(:disabled) { background: var(--color-primary-hover); }

.btn-danger { background: rgba(239,68,68,0.1); color: var(--color-danger); }
.btn-danger:hover:not(:disabled) { background: var(--color-danger); color: white; }

.btn-outline { background: transparent; border-color: var(--border-color); color: var(--text-main); }
.btn-outline:hover:not(:disabled) { border-color: var(--text-muted); }
</style>