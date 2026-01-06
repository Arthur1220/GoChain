<script setup>
import { onMounted, ref } from 'vue';
import { useContractStore } from '../stores/contractStore';
import BaseCard from '../components/ui/BaseCard.vue';
import BaseInput from '../components/ui/BaseInput.vue';
import BaseButton from '../components/ui/BaseButton.vue';
import TokenCard from '../components/domain/tokens/TokenCard.vue';
import { ShieldAlert } from 'lucide-vue-next';

const store = useContractStore();
const form = ref({ address: '', name: '', symbol: '', decimals: 18 });

const handleAdd = async () => {
  await store.addContract(form.value);
  form.value = { address: '', name: '', symbol: '', decimals: 18 }; // Reset
};

onMounted(() => store.fetchContracts());
</script>

<template>
  <div class="admin-view">
    <div class="header">
      <h2><ShieldAlert class="text-primary"/> Área Administrativa</h2>
    </div>

    <div class="content-grid">
      <BaseCard>
        <h3 class="card-title">Novo Monitoramento</h3>
        <form @submit.prevent="handleAdd" class="form-stack">
          <BaseInput v-model="form.name" placeholder="Nome (Ex: Wrapped Ether)" />
          <div class="row">
            <BaseInput v-model="form.symbol" placeholder="Símbolo (Ex: WETH)" />
            <BaseInput v-model.number="form.decimals" type="number" placeholder="Decimais (18)" />
          </div>
          <BaseInput v-model="form.address" placeholder="Endereço (0x...)" class="font-mono" />
          
          <BaseButton :loading="store.loading" class="w-full">
            Adicionar Token
          </BaseButton>
          
          <p v-if="store.error" class="text-danger text-sm">{{ store.error }}</p>
        </form>
      </BaseCard>

      <div class="tokens-column">
        <h3 class="column-title">Tokens Ativos</h3>
        <div class="tokens-list">
           <template v-for="c in store.contracts" :key="c?.address">
             <TokenCard 
               v-if="c"
               :token="c" 
               :isAdmin="true" 
               @remove="store.removeContract(c.address)"
             />
           </template>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.header { margin-bottom: 2rem; display: flex; align-items: center; gap: 10px; }
.content-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 2rem; }
.form-stack { display: flex; flex-direction: column; gap: 1rem; margin-top: 1.5rem; }
.row { display: flex; gap: 1rem; }
.w-full { width: 100%; justify-content: center; }
.card-title { font-size: 1.1rem; font-weight: 600; color: var(--text-main); }
.column-title { font-size: 1.1rem; margin-bottom: 1rem; color: var(--text-muted); }
.tokens-list { display: flex; flex-direction: column; gap: 1rem; }
.text-sm { font-size: 0.85rem; margin-top: 0.5rem; }
</style>