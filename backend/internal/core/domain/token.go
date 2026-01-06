package domain

import "time"

// Token representa um contrato ERC-20 que estamos monitorando
type Token struct {
	Address   string    `json:"address"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	Decimals  int       `json:"decimals"` // Essencial para calcular o valor correto
	CreatedAt time.Time `json:"created_at"`
}
