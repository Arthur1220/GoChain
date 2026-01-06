import { defineStore } from 'pinia';
import api from '../services/api';
import { ref, computed } from 'vue';

export const useTransactionStore = defineStore('transaction', () => {
  // --- STATE ---
  const transfers = ref([]);
  const stats = ref({
    total_volume: 0,
    max_whale: 0,
    total_count: 0,
    first_seen_date: null
  });
  
  const selectedContract = ref(null);
  const loading = ref(false);
  const error = ref(null);
  
  // Paginação e Filtros
  const currentPage = ref(1);
  const itemsPerPage = 10;
  const searchQuery = ref('');
  const sortOrder = ref('newest');

  // Controle de Polling (Loop)
  let pollingInterval = null;

  // --- ACTIONS ---

  // 1. Define qual contrato estamos olhando (Limpa dados antigos imediatamente)
  const setContract = (address) => {
    if (selectedContract.value !== address) {
      selectedContract.value = address;
      currentPage.value = 1;
      transfers.value = []; // Limpa visualmente para forçar loading
      searchQuery.value = ''; // Reseta busca
      
      // Busca dados novos imediatamente
      fetchData();
    }
  };

  // 2. Busca Dados (Transações + Stats)
  const fetchData = async () => {
    if (!selectedContract.value) return;

    const isFirstLoad = transfers.value.length === 0;
    if (isFirstLoad) loading.value = true;

    try {
      const params = {
        page: currentPage.value,
        limit: itemsPerPage,
        contract: selectedContract.value
      };

      const [txRes, statsRes] = await Promise.all([
        api.get('/transfers', { params }),
        api.get('/stats', { params: { contract: selectedContract.value } })
      ]);

      // CORREÇÃO AQUI TAMBÉM: Acessando .data.data
      transfers.value = txRes.data?.data || [];
      
      // Stats também vem dentro de .data
      stats.value = statsRes.data?.data || stats.value;
      
      error.value = null;

    } catch (err) {
      console.error("Erro no fetch:", err);
      if (isFirstLoad) error.value = "Falha ao sincronizar com a Blockchain.";
    } finally {
      loading.value = false;
    }
  };

  // 3. Controle do Polling (Automático)
  const startPolling = () => {
    if (pollingInterval) clearInterval(pollingInterval);
    fetchData(); // Chama uma vez agora
    pollingInterval = setInterval(fetchData, 5000); // Chama a cada 5s
  };

  const stopPolling = () => {
    if (pollingInterval) clearInterval(pollingInterval);
    pollingInterval = null;
  };

  // 4. Paginação
  const changePage = (page) => {
    currentPage.value = page;
    fetchData();
  };

  // --- GETTERS (Dados Computados / Tratados) ---

  const processedTransfers = computed(() => {
    let data = [...transfers.value];

    // Filtro Local (Client-Side Search na página atual)
    if (searchQuery.value) {
      const q = searchQuery.value.toLowerCase();
      data = data.filter(tx => 
        tx.from.toLowerCase().includes(q) || 
        tx.to.toLowerCase().includes(q) || 
        tx.tx_hash.toLowerCase().includes(q)
      );
    }

    // Ordenação Local
    if (sortOrder.value === 'value-desc') {
      data.sort((a, b) => b.amount - a.amount);
    } else if (sortOrder.value === 'value-asc') {
      data.sort((a, b) => a.amount - b.amount);
    }

    return data;
  });

  const totalPages = computed(() => {
    // Calcula páginas com base no total vindo do banco (stats.total_count)
    return Math.ceil((stats.value.total_count || 0) / itemsPerPage) || 1;
  });

  return {
    // State
    transfers,
    stats,
    selectedContract,
    loading,
    error,
    searchQuery,
    sortOrder,
    currentPage,
    
    // Getters
    processedTransfers,
    totalPages,

    // Actions
    setContract,
    fetchData,
    startPolling,
    stopPolling,
    changePage
  };
});