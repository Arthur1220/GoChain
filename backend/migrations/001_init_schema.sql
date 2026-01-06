-- Habilita extensão para gerar UUIDs (opcional, mas boa prática)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabela de Contratos Monitorados
CREATE TABLE IF NOT EXISTS watched_contracts (
    address VARCHAR(42) PRIMARY KEY,
    name VARCHAR(100),
    symbol VARCHAR(20),
    decimals INT DEFAULT 18,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de Controle de Blocos (Checkpoints)
CREATE TABLE IF NOT EXISTS processed_blocks (
    contract_address VARCHAR(42) PRIMARY KEY,
    last_block BIGINT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de Transferências (AQUI ESTÁ A CORREÇÃO DO DECIMAL)
CREATE TABLE IF NOT EXISTS usdc_transfers (
    id SERIAL PRIMARY KEY,
    tx_hash VARCHAR(66) NOT NULL,
    from_addr VARCHAR(42) NOT NULL,
    to_addr VARCHAR(42) NOT NULL,
    
    -- NUMERIC(80, 18) suporta números astronômicos de tokens com 18 casas decimais
    amount NUMERIC(80, 18) NOT NULL, 
    
    block_number BIGINT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    
    -- Restrição de Unicidade para evitar duplicatas
    CONSTRAINT unique_tx_event UNIQUE (tx_hash, from_addr, to_addr, amount) 
);

-- Índices para performance
CREATE INDEX IF NOT EXISTS idx_transfers_contract ON usdc_transfers(contract_address);
CREATE INDEX IF NOT EXISTS idx_transfers_block ON usdc_transfers(block_number DESC);