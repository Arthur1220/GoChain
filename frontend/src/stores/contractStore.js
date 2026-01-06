import { defineStore } from 'pinia';
import api from '../services/api';
import { ref } from 'vue';

export const useContractStore = defineStore('contract', () => {
  const contracts = ref([]);
  const loading = ref(false);
  const error = ref(null);

  // Busca a lista de contratos (Home e Admin)
  const fetchContracts = async () => {
    loading.value = true;
    error.value = null;
    try {
      const res = await api.get('/contracts');
      
      // CORREÇÃO CRÍTICA AQUI:
      // res.data é o corpo da resposta HTTP
      // res.data.data é onde o nosso Backend Go colocou a lista
      const rawData = res.data?.data || []; 
      
      if (Array.isArray(rawData)) {
        contracts.value = rawData.filter(c => c && c.address);
      } else {
        contracts.value = [];
        console.warn("Formato de resposta inesperado:", rawData);
      }

    } catch (err) {
      console.error(err);
      error.value = 'Falha ao carregar contratos.';
    } finally {
      loading.value = false;
    }
  };

  // Adiciona novo contrato (Admin)
  const addContract = async (form) => {
    loading.value = true;
    try {
      await api.post('/contracts', form);
      await fetchContracts(); // Atualiza a lista local
      return true; // Sucesso
    } catch (err) {
      throw err; // Joga erro pra View tratar msg
    } finally {
      loading.value = false;
    }
  };

  // Remove contrato (Admin)
  const removeContract = async (address) => {
    try {
      await api.delete(`/contracts/${address}`);
      contracts.value = contracts.value.filter(c => c.address !== address);
    } catch (err) {
      console.error(err);
      throw new Error('Erro ao remover contrato');
    }
  };

  return {
    contracts,
    loading,
    error,
    fetchContracts,
    addContract,
    removeContract
  };
});