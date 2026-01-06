<script setup>
import { ExternalLink } from 'lucide-vue-next';
import { formatCurrency, formatDate, shortHash } from '../../../utils/formatters';

defineProps({
  transfers: { type: Array, required: true },
  highlight: { type: String, default: '' }
});

// Função para abrir Etherscan
const openExplorer = (hash) => {
  window.open(`https://etherscan.io/tx/${hash}`, '_blank');
};
</script>

<template>
  <div class="table-wrapper">
    <table>
      <thead>
        <tr>
          <th>TX Hash</th>
          <th>Tempo</th>
          <th>De (From)</th>
          <th>Para (To)</th>
          <th class="text-right">Valor</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="tx in transfers" :key="tx.id">
          <td class="hash-col">
            <span class="font-mono text-muted">{{ shortHash(tx.tx_hash) }}</span>
            <button class="icon-btn" @click.stop="openExplorer(tx.tx_hash)" title="Ver no Etherscan">
              <ExternalLink :size="12" />
            </button>
          </td>

          <td class="date-col">{{ formatDate(tx.timestamp) }}</td>

          <td class="addr-col font-mono">
            <span :class="{ 'highlight-bg': highlight && tx.from.toLowerCase().includes(highlight.toLowerCase()) }">
              {{ shortHash(tx.from) }}
            </span>
          </td>
          <td class="addr-col font-mono">
            <span :class="{ 'highlight-bg': highlight && tx.to.toLowerCase().includes(highlight.toLowerCase()) }">
              {{ shortHash(tx.to) }}
            </span>
          </td>

          <td class="text-right amount-col font-mono text-primary">
            {{ formatCurrency(tx.amount) }}
          </td>
        </tr>
      </tbody>
    </table>

    <div v-if="transfers.length === 0" class="empty-state">
      Nenhuma transação encontrada.
    </div>
  </div>
</template>

<style scoped>
.table-wrapper {
  overflow-x: auto;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  background: var(--bg-card);
}

table { width: 100%; border-collapse: collapse; font-size: 0.9rem; }
th { text-align: left; padding: 1rem; color: var(--text-muted); font-weight: 600; border-bottom: 1px solid var(--border-color); background: rgba(0,0,0,0.2); }
td { padding: 1rem; border-bottom: 1px solid rgba(255,255,255,0.03); color: var(--text-main); }
tr:last-child td { border-bottom: none; }
tr:hover td { background: rgba(255,255,255,0.02); }

.text-right { text-align: right; }
.hash-col { display: flex; align-items: center; gap: 8px; }
.icon-btn { background: none; border: none; color: var(--text-muted); cursor: pointer; opacity: 0.5; padding: 0; display: flex; }
.icon-btn:hover { opacity: 1; color: var(--color-primary); }

.highlight-bg { background: rgba(34, 197, 94, 0.2); color: white; padding: 2px 4px; border-radius: 4px; }
.empty-state { padding: 3rem; text-align: center; color: var(--text-muted); }
</style>