package postgres

import (
	"database/sql"
	"go-chain/internal/core/domain"
)

func (r *PostgresRepository) AddToken(token *domain.Token) error {
	query := `
		INSERT INTO watched_contracts (address, name, symbol, decimals, created_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (address) DO NOTHING`

	_, err := r.db.Exec(query, token.Address, token.Name, token.Symbol, token.Decimals, token.CreatedAt)
	return err
}

func (r *PostgresRepository) RemoveToken(address string) error {
	// Primeiro removemos o checkpoint para não ficar lixo
	_, _ = r.db.Exec("DELETE FROM processed_blocks WHERE contract_address = $1", address)

	// Depois removemos o contrato
	_, err := r.db.Exec("DELETE FROM watched_contracts WHERE address = $1", address)
	return err
}

func (r *PostgresRepository) ListTokens() ([]*domain.Token, error) {
	rows, err := r.db.Query("SELECT address, name, symbol, decimals, created_at FROM watched_contracts ORDER BY symbol")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []*domain.Token
	for rows.Next() {
		var t domain.Token
		if err := rows.Scan(&t.Address, &t.Name, &t.Symbol, &t.Decimals, &t.CreatedAt); err != nil {
			return nil, err
		}
		tokens = append(tokens, &t)
	}
	return tokens, nil
}

func (r *PostgresRepository) GetToken(address string) (*domain.Token, error) {
	var t domain.Token
	err := r.db.QueryRow("SELECT address, name, symbol, decimals, created_at FROM watched_contracts WHERE address = $1", address).
		Scan(&t.Address, &t.Name, &t.Symbol, &t.Decimals, &t.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &t, err
}
