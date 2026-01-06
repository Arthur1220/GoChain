package domain

import "time"

// Transfer representa uma movimentação na blockchain já processada
type Transfer struct {
	ID              int64     `json:"id"`
	TxHash          string    `json:"tx_hash"`
	From            string    `json:"from"`
	To              string    `json:"to"`
	Amount          float64   `json:"amount"` // Valor já dividido pelos decimais (human readable)
	BlockNumber     uint64    `json:"block_number"`
	Timestamp       time.Time `json:"timestamp"`
	ContractAddress string    `json:"contract_address"`
}

// TokenStats agrupa os dados para os Cards do Dashboard
type TokenStats struct {
	TotalVolume   float64   `json:"total_volume"`
	MaxWhale      float64   `json:"max_whale"`
	TotalCount    int64     `json:"total_count"`
	FirstSeenDate time.Time `json:"first_seen_date"`
}
