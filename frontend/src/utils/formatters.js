// Formata valores monetários (Ex: $1,200.00 ou $1.5M)
export const formatCurrency = (value) => {
  if (value === undefined || value === null) return '$0';
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
    maximumFractionDigits: 2,
    notation: value > 10000 ? 'compact' : 'standard' 
  }).format(value);
};

// Formata data (Ex: 05/01/2026)
export const formatDate = (dateString) => {
  if (!dateString || dateString.startsWith('0001')) return 'N/A';
  return new Date(dateString).toLocaleDateString('pt-BR', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  });
};

// Encurta Hash (Ex: 0x1234...5678)
export const shortHash = (hash) => {
  if (!hash) return '';
  return `${hash.slice(0, 6)}...${hash.slice(-4)}`;
};