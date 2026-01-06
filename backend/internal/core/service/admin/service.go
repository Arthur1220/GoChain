package admin

import (
	"go-chain/internal/core/domain"
	"go-chain/internal/core/ports"
	"time"
)

type AdminService struct {
	tokenRepo ports.TokenRepository
	txRepo    ports.TransferRepository
}

func NewAdminService(tRepo ports.TokenRepository, txRepo ports.TransferRepository) *AdminService {
	return &AdminService{
		tokenRepo: tRepo,
		txRepo:    txRepo,
	}
}

// --- Gerenciamento de Tokens ---

func (s *AdminService) AddToken(addr, name, symbol string, decimals int) error {
	token := &domain.Token{
		Address:   addr,
		Name:      name,
		Symbol:    symbol,
		Decimals:  decimals,
		CreatedAt: time.Now(),
	}
	return s.tokenRepo.AddToken(token)
}

func (s *AdminService) RemoveToken(addr string) error {
	return s.tokenRepo.RemoveToken(addr)
}

func (s *AdminService) ListTokens() ([]*domain.Token, error) {
	return s.tokenRepo.ListTokens()
}

// --- Dados para o Dashboard ---

func (s *AdminService) GetDashboardData(contractAddr string, page, limit int) ([]*domain.Transfer, *domain.TokenStats, error) {
	// Busca lista paginada
	transfers, err := s.txRepo.ListTransfers(contractAddr, page, limit)
	if err != nil {
		return nil, nil, err
	}

	// Busca estatísticas globais
	stats, err := s.txRepo.GetTokenStats(contractAddr)
	if err != nil {
		return nil, nil, err
	}

	return transfers, stats, nil
}
