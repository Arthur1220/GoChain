<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import { useContractStore } from '../stores/contractStore';
import api from '../services/api';
import TokenCard from '../components/domain/tokens/TokenCard.vue';
import TheSpinner from '../components/ui/TheSpinner.vue';

const router = useRouter();
const contractStore = useContractStore();
const enrichedContracts = ref([]);
let polling = null;

// Lógica de buscar estatísticas para os cards
const fetchStats = async () => {
  if (contractStore.contracts.length === 0) await contractStore.fetchContracts();
  
  const promises = contractStore.contracts.map(async (token) => {
    try {
      const { data } = await api.get('/stats', { params: { contract: token.address } });
      // Lembra da correção do .data.data? Aqui api.get já retorna o response, 
      // então data.data é onde está o payload se usar response.go
      return { ...token, stats: data.data || null };
    } catch {
      return { ...token, stats: null };
    }
  });
  enrichedContracts.value = await Promise.all(promises);
};

onMounted(() => {
  fetchStats();
  polling = setInterval(fetchStats, 10000);
});

onUnmounted(() => { if (polling) clearInterval(polling); });

const goToDash = (addr) => router.push({ name: 'dashboard', query: { contract: addr } });
</script>

<template>
  <div class="home-view">
    <div class="hero">
      <h1>Go-Chain <span class="text-primary">Monitor</span></h1>
      <p class="text-muted">Inteligência Blockchain em Tempo Real</p>
    </div>

    <div v-if="contractStore.loading && enrichedContracts.length === 0" class="center-state">
      <TheSpinner />
    </div>

    <div v-else-if="enrichedContracts.length === 0" class="center-state text-muted">
      Nenhum ativo configurado. Vá ao Admin.
    </div>

    <div v-else class="grid">
      <TokenCard 
        v-for="token in enrichedContracts" 
        :key="token.address" 
        :token="token"
        @click="goToDash(token.address)"
      />
    </div>
  </div>
</template>

<style scoped>
.hero { text-align: center; margin-bottom: 4rem; }
.hero h1 { font-size: 3rem; font-weight: 800; letter-spacing: -2px; margin-bottom: 0.5rem; }
.grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(300px, 1fr)); gap: 1.5rem; }
.center-state { display: flex; justify-content: center; padding: 4rem; }
</style>