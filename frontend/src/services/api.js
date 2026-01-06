import axios from 'axios';

// Configura a URL base do seu Backend Go
const api = axios.create({
    baseURL: 'http://localhost:8080/api/v1',
    timeout: 5000, // 5 segundos de timeout
});

export default api;