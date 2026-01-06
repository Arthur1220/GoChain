package ports

import "go-chain/internal/core/domain"

// TokenRepository define as operações com os contratos monitorados
type TokenRepository interface {
	AddToken(token *domain.Token) error
	RemoveToken(address string) error
	ListTokens() ([]*domain.Token, error)
	GetToken(address string) (*domain.Token, error)
}

// TransferRepository define as operações com as transações
type TransferRepository interface {
	// Salva uma nova transferência
	SaveTransfer(tx *domain.Transfer) error

	// Controle de qual bloco paramos
	GetLastProcessedBlock(address string) (uint64, error)
	UpdateLastProcessedBlock(address string, block uint64) error

	// Consultas para o Frontend
	ListTransfers(contractAddress string, page, limit int) ([]*domain.Transfer, error)
	GetTokenStats(contractAddress string) (*domain.TokenStats, error)
}
