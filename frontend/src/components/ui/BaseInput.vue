<script setup>
defineProps(['modelValue', 'placeholder', 'type']);
defineEmits(['update:modelValue']);
</script>

<template>
  <input 
    class="base-input"
    :type="type || 'text'"
    :placeholder="placeholder"
    :value="modelValue"
    @input="$emit('update:modelValue', $event.target.value)"
  />
</template>

<style scoped>
.base-input {
  background: var(--bg-input);
  border: 1px solid var(--border-color);
  color: var(--text-main);
  padding: 0.75rem 1rem;
  border-radius: var(--radius-md);
  width: 100%;
  font-family: inherit;
  outline: none;
  transition: border-color 0.2s;
}
.base-input:focus { border-color: var(--color-primary); }
</style>