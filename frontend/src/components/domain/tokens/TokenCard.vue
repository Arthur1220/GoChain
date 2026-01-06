<script setup>
import { Activity, Calendar, ArrowRight, Trash2 } from 'lucide-vue-next';
import BaseCard from '../../ui/BaseCard.vue';
import { formatCurrency, formatDate } from '../../../utils/formatters';

const props = defineProps({
  token: Object,
  isAdmin: Boolean // Se true, mostra botão de deletar. Se false, mostra seta de Dashboard.
});

const emit = defineEmits(['click', 'remove']);
</script>

<template>
  <BaseCard 
    class="token-card" 
    :class="{ 'clickable': !isAdmin }"
    @click="!isAdmin && emit('click')"
  >
    <div class="card-header">
      <div class="token-identity">
        <span class="symbol">{{ token.symbol }}</span>
        <span class="name">{{ token.name }}</span>
      </div>
      <button 
        v-if="isAdmin" 
        class="btn-icon-danger"
        @click.stop="emit('remove')"
      >
        <Trash2 :size="18" />
      </button>
    </div>

    <div class="card-stats" v-if="token.stats">
      <div class="stat-row">
        <span class="label"><Activity :size="14"/> Volume</span>
        <span class="value highlight">{{ formatCurrency(token.stats.total_volume) }}</span>
      </div>
      <div class="stat-row">
        <span class="label"><Calendar :size="14"/> Início</span>
        <span class="value">{{ formatDate(token.stats.first_seen_date) }}</span>
      </div>
    </div>
    
    <div class="card-info" v-else>
      <span class="address font-mono">{{ token.address }}</span>
    </div>

    <div class="card-footer" v-if="!isAdmin">
      <span class="action-text">Ver Dashboard</span>
      <ArrowRight :size="16" />
    </div>
  </BaseCard>
</template>

<style scoped>
.token-card {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 180px;
  transition: transform 0.2s, border-color 0.2s;
}
.token-card.clickable { cursor: pointer; }
.token-card.clickable:hover {
  transform: translateY(-4px);
  border-color: var(--color-primary);
}

.card-header { display: flex; justify-content: space-between; align-items: flex-start; margin-bottom: 1rem; }
.token-identity { display: flex; flex-direction: column; }
.symbol { font-size: 1.5rem; font-weight: 700; font-family: var(--font-mono); color: var(--text-main); }
.name { font-size: 0.85rem; color: var(--text-muted); text-transform: uppercase; letter-spacing: 1px; }

.stat-row { display: flex; justify-content: space-between; margin-bottom: 0.5rem; font-size: 0.9rem; }
.label { display: flex; align-items: center; gap: 6px; color: var(--text-muted); }
.highlight { color: var(--color-primary); font-family: var(--font-mono); font-weight: 600; }

.address { font-size: 0.75rem; color: var(--text-muted); word-break: break-all; opacity: 0.6; }

.card-footer { margin-top: auto; padding-top: 1rem; border-top: 1px solid rgba(255,255,255,0.05); display: flex; justify-content: space-between; align-items: center; font-size: 0.85rem; color: var(--text-muted); }
.token-card:hover .action-text { color: var(--color-primary); }

.btn-icon-danger {
  background: transparent; border: none; color: var(--color-danger);
  cursor: pointer; opacity: 0.7; transition: 0.2s;
}
.btn-icon-danger:hover { opacity: 1; transform: scale(1.1); }
</style>