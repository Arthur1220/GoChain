import { createApp } from 'vue'
import { createPinia } from 'pinia' // <--- Importar
import './assets/main.css'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia()) // <--- Ativar antes do router
app.use(router)

app.mount('#app')