# 🚀 Go-Chain Monitor

> **Inteligência Blockchain em Tempo Real.** Uma plataforma Full-Stack de alta performance para monitoramento, indexação e análise de transações ERC-20 na rede Ethereum.

<div align="center">

![Status](https://img.shields.io/badge/STATUS-OPERATIONAL-success?style=for-the-badge)
![License](https://img.shields.io/badge/LICENSE-MIT-green?style=for-the-badge)

<br>

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Ethereum](https://img.shields.io/badge/Ethereum-3C3C3D?style=for-the-badge&logo=Ethereum&logoColor=white)

<br>

![Vue.js](https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D)
![Pinia](https://img.shields.io/badge/pinia-%23F7D336.svg?style=for-the-badge&logo=pinia&logoColor=000)
![Vite](https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white)

</div>

---

## 🎯 Sobre o Projeto

O **Go-Chain Monitor** resolve um problema crítico no ecossistema cripto: a necessidade de rastrear grandes volumes de transações com precisão matemática absoluta e baixa latência.

Diferente de exploradores genéricos, este sistema foca em **tokens específicos** (como USDT, WETH), permitindo uma visão granular de "Baleias" (Whales), volume transacionado e métricas em tempo real.

### 🌟 Destaques de Engenharia

* **Precisão Financeira:** Utiliza tipos de dados `NUMERIC(80, 18)` no PostgreSQL e `big.Int/big.Float` no Go para garantir que **nenhum wei** (a menor unidade do Ethereum) seja perdido em arredondamentos.
* **Arquitetura Hexagonal:** O Backend é desacoplado, facilitando testes e troca de tecnologias (ex: mudar de Postgres para Mongo sem tocar na regra de negócio).
* **Frontend Atômico:** Interface construída com **Atomic Design** e **CSS Variables**, garantindo consistência visual e fácil manutenção sem dependência de frameworks CSS pesados.
* **Self-Healing:** O sistema detecta falhas na estrutura do banco de dados e aplica migrações automáticas ao iniciar (`Auto-Migration`).

---

## 📸 Screenshots

*(Espaço reservado para as imagens que você gerou. Sugiro colocar um GIF ou print do Dashboard aqui)*

| Dashboard | Admin Panel |
| --- | --- |
| *Visualização de métricas e tabela em tempo real* | *Gerenciamento de tokens monitorados* |

---

## 🛠️ Tech Stack

### 🟢 Backend (The Engine)

Construído para concorrência e velocidade.

* **Linguagem:** Go (Golang)
* **Framework Web:** Gin Gonic (API RESTful)
* **Blockchain Client:** Go-Ethereum (go-eth) client RPC
* **Database Driver:** `lib/pq` (Conexão nativa e performática)
* **Arquitetura:** Clean Architecture (Ports & Adapters)

### 🔵 Frontend (The Experience)

Reativo, moderno e organizado.

* **Framework:** Vue.js 3 (Composition API)
* **State Management:** Pinia (Stores modulares para Contratos e Transações)
* **Build Tool:** Vite
* **Estilização:** CSS Scoped com Design System (Variáveis CSS)
* **Arquitetura:** Atomic Design (Atoms, Molecules, Organisms)

### 🟣 Database (The Vault)

* **SGBD:** PostgreSQL
* **Features:** Índices otimizados para busca por hash e bloco.

---

## 🏗️ Arquitetura do Sistema

### Backend: Clean Architecture

O projeto segue estritamente a separação de responsabilidades:

1. **Core (Domain & Ports):** Define as entidades (`Token`, `Transfer`) e as interfaces (`Repository`, `BlockchainService`). Não sabe que banco ou API existe.
2. **Services:** A regra de negócio. Onde ocorre a matemática de conversão de decimais e a orquestração.
3. **Adapters:**
* **EthClient:** Implementa a comunicação com a Blockchain.
* **PostgresRepo:** Implementa a persistência dos dados.
* **HTTP Handler:** Expõe os dados para o mundo via JSON.



### Frontend: Pinia + Atomic Design

O estado é gerenciado globalmente para evitar "Prop Drilling":

* **ContractStore:** Gerencia a lista de ativos (CRUD).
* **TransactionStore:** Gerencia o fluxo pesado de dados, paginação e *polling* automático.

A UI é construída com componentes base (`BaseCard`, `BaseButton`) que são compostos para criar interfaces complexas (`TransactionTable`).

---

## 🚀 Como Rodar Localmente

### Pré-requisitos

* Go 1.21+
* Node.js 18+
* PostgreSQL
* Uma URL RPC da Ethereum (ex: Infura, Alchemy ou LlamaNodes grátis).

### 1. Configuração do Banco de Dados

Crie um banco de dados no Postgres chamado `gochain`. O sistema criará as tabelas automaticamente.

### 2. Backend

```bash
cd backend

# Configure as variáveis (Opcional, o sistema tem defaults)
export DATABASE_URL="postgres://user:senha@localhost:5432/gochain?sslmode=disable"
export RPC_URL="https://eth.llamarpc.com"

# Instale dependências
go mod tidy

# Rode o servidor
go run cmd/api/main.go

```

*O servidor iniciará em `localhost:8080` e iniciará o worker de monitoramento.*

### 3. Frontend

```bash
cd frontend

# Instale dependências
npm install

# Rode o servidor de desenvolvimento
npm run dev

```

*Acesse `localhost:5173` no seu navegador.*

---

## 🤝 Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir Issues ou Pull Requests.

---

Desenvolvido com 💚 por **Arthur Marques Azevedo**