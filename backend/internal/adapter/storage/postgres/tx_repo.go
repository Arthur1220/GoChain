package postgres

import (
	"database/sql"
	"go-chain/internal/core/domain"
)

// --- CHECKPOINTS (Controle de Blocos) ---

func (r *PostgresRepository) GetLastProcessedBlock(address string) (uint64, error) {
	var block uint64
	err := r.db.QueryRow("SELECT last_block FROM processed_blocks WHERE contract_address = $1", address).Scan(&block)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return block, err
}

func (r *PostgresRepository) UpdateLastProcessedBlock(address string, block uint64) error {
	query := `
		INSERT INTO processed_blocks (contract_address, last_block, updated_at)
		VALUES ($1, $2, NOW())
		ON CONFLICT (contract_address) 
		DO UPDATE SET last_block = $2, updated_at = NOW()`
	_, err := r.db.Exec(query, address, block)
	return err
}

// --- TRANSAÇÕES ---

func (r *PostgresRepository) SaveTransfer(tx *domain.Transfer) error {
	query := `
		INSERT INTO usdc_transfers (tx_hash, from_addr, to_addr, amount, block_number, timestamp, contract_address)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT ON CONSTRAINT unique_tx_event DO NOTHING`

	_, err := r.db.Exec(query, tx.TxHash, tx.From, tx.To, tx.Amount, tx.BlockNumber, tx.Timestamp, tx.ContractAddress)
	return err
}

func (r *PostgresRepository) ListTransfers(contractAddress string, page, limit int) ([]*domain.Transfer, error) {
	offset := (page - 1) * limit
	query := `
		SELECT id, tx_hash, from_addr, to_addr, amount, block_number, timestamp, contract_address
		FROM usdc_transfers
		WHERE contract_address = $1
		ORDER BY block_number DESC, id DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(query, contractAddress, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []*domain.Transfer
	for rows.Next() {
		var tx domain.Transfer
		if err := rows.Scan(&tx.ID, &tx.TxHash, &tx.From, &tx.To, &tx.Amount, &tx.BlockNumber, &tx.Timestamp, &tx.ContractAddress); err != nil {
			return nil, err
		}
		transfers = append(transfers, &tx)
	}
	return transfers, nil
}

// --- ESTATÍSTICAS ---

func (r *PostgresRepository) GetTokenStats(contractAddress string) (*domain.TokenStats, error) {
	query := `
		SELECT 
			COALESCE(SUM(amount), 0) as total_vol,
			COALESCE(MAX(amount), 0) as max_whale,
			COUNT(*) as total_count,
			COALESCE(MIN(timestamp), '0001-01-01') as first_seen
		FROM usdc_transfers
		WHERE contract_address = $1`

	var stats domain.TokenStats
	err := r.db.QueryRow(query, contractAddress).
		Scan(&stats.TotalVolume, &stats.MaxWhale, &stats.TotalCount, &stats.FirstSeenDate)

	return &stats, err
}
