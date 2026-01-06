package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq" // Driver do Postgres
)

// PostgresRepository implementa as interfaces TokenRepository e TransferRepository
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository cria uma nova conexão
func NewPostgresRepository(connStr string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao pingar banco: %w", err)
	}

	repo := &PostgresRepository{db: db}

	// --- AUTO-MIGRAÇÃO ---
	// Executa a criação das tabelas ao iniciar
	if err := repo.initSchema(); err != nil {
		return nil, fmt.Errorf("erro ao criar tabelas: %w", err)
	}

	return repo, nil
}

// Close fecha a conexão (útil para graceful shutdown)
func (r *PostgresRepository) Close() error {
	return r.db.Close()
}

// initSchema contém o SQL para garantir que as tabelas existam
func (r *PostgresRepository) initSchema() error {
	slog.Info("Verificando esquema do banco de dados...")

	query := `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	CREATE TABLE IF NOT EXISTS watched_contracts (
		address VARCHAR(42) PRIMARY KEY,
		name VARCHAR(100),
		symbol VARCHAR(20),
		decimals INT DEFAULT 18,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS processed_blocks (
		contract_address VARCHAR(42) PRIMARY KEY,
		last_block BIGINT NOT NULL,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS usdc_transfers (
		id SERIAL PRIMARY KEY,
		tx_hash VARCHAR(66) NOT NULL,
		from_addr VARCHAR(42) NOT NULL,
		to_addr VARCHAR(42) NOT NULL,
		amount NUMERIC(80, 18) NOT NULL, 
		block_number BIGINT NOT NULL,
		timestamp TIMESTAMP NOT NULL,
		contract_address VARCHAR(42) NOT NULL,
		CONSTRAINT unique_tx_event UNIQUE (tx_hash, from_addr, to_addr, amount) 
	);

	CREATE INDEX IF NOT EXISTS idx_transfers_contract ON usdc_transfers(contract_address);
	CREATE INDEX IF NOT EXISTS idx_transfers_block ON usdc_transfers(block_number DESC);
	`

	_, err := r.db.Exec(query)
	return err
}
