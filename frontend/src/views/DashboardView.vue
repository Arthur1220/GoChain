<script setup>
import { onMounted, onUnmounted } from 'vue';
import { useRoute } from 'vue-router';
import { useTransactionStore } from '../stores/transactionStore';
import { useContractStore } from '../stores/contractStore';
import { formatCurrency } from '../utils/formatters';

import BaseCard from '../components/ui/BaseCard.vue';
import TransactionTable from '../components/domain/transactions/TransactionTable.vue';
import TheSpinner from '../components/ui/TheSpinner.vue';

const route = useRoute();
const txStore = useTransactionStore();
const contractStore = useContractStore();

onMounted(async () => {
  if (contractStore.contracts.length === 0) await contractStore.fetchContracts();
  
  const initial = route.query.contract || contractStore.contracts[0]?.address;
  if (initial) txStore.setContract(initial);
  
  txStore.startPolling();
});

onUnmounted(() => txStore.stopPolling());
</script>

<template>
  <div class="dashboard">
    <div class="toolbar">
      <select 
        class="contract-select"
        :value="txStore.selectedContract" 
        @change="e => txStore.setContract(e.target.value)"
      >
        <option v-for="c in contractStore.contracts" :key="c.address" :value="c.address">
          {{ c.symbol }} - {{ c.name }}
        </option>
      </select>
      
      <span v-if="txStore.loading && txStore.transfers.length > 0" class="live-badge">
        <span class="dot"></span> LIVE
      </span>
    </div>

    <div class="stats-grid">
      <BaseCard>
        <span class="stat-label">Volume Total</span>
        <div class="stat-value text-primary">{{ formatCurrency(txStore.stats.total_volume) }}</div>
      </BaseCard>
      <BaseCard>
        <span class="stat-label">Maior Baleia</span>
        <div class="stat-value">{{ formatCurrency(txStore.stats.max_whale) }}</div>
      </BaseCard>
      <BaseCard>
        <span class="stat-label">Transações</span>
        <div class="stat-value">{{ txStore.stats.total_count }}</div>
      </BaseCard>
    </div>

    <div class="table-section">
      <div v-if="txStore.loading && txStore.transfers.length === 0" class="loading-box">
        <TheSpinner />
        <p>Sincronizando Blockchain...</p>
      </div>

      <TransactionTable 
        v-else 
        :transfers="txStore.processedTransfers" 
      />
    </div>
  </div>
</template>

<style scoped>
.dashboard { display: flex; flex-direction: column; gap: 2rem; }

.toolbar { display: flex; justify-content: space-between; align-items: center; }
.contract-select { 
  background: var(--bg-input); color: var(--text-main); border: 1px solid var(--border-color);
  padding: 0.5rem 1rem; border-radius: var(--radius-md); font-family: var(--font-sans);
}

.stats-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 1.5rem; }
.stat-label { font-size: 0.85rem; color: var(--text-muted); text-transform: uppercase; letter-spacing: 1px; }
.stat-value { font-size: 1.8rem; font-weight: 700; font-family: var(--font-mono); margin-top: 0.5rem; }

.loading-box { display: flex; flex-direction: column; align-items: center; gap: 1rem; padding: 4rem; color: var(--text-muted); }

.live-badge { display: flex; align-items: center; gap: 6px; font-size: 0.75rem; color: var(--color-primary); font-weight: 700; }
.dot { width: 6px; height: 6px; background: currentColor; border-radius: 50%; animation: pulse 1s infinite; }
@keyframes pulse { 50% { opacity: 0.5; } }
</style>