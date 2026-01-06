package monitor

import (
	"log/slog"
	"strings"
	"time"

	"go-chain/internal/core/domain"
	"go-chain/internal/core/ports"

	"github.com/ethereum/go-ethereum/common"
)

type MonitorService struct {
	ethClient ports.BlockchainService
	tokenRepo ports.TokenRepository
	txRepo    ports.TransferRepository
}

func NewMonitorService(eth ports.BlockchainService, tRepo ports.TokenRepository, txRepo ports.TransferRepository) *MonitorService {
	return &MonitorService{
		ethClient: eth,
		tokenRepo: tRepo,
		txRepo:    txRepo,
	}
}

// ProcessRange busca logs de todos os tokens monitorados e salva as transferências
func (s *MonitorService) ProcessRange(fromBlock, toBlock int64) error {
	// 1. Busca quais tokens estamos monitorando (para saber endereços e decimais)
	tokens, err := s.tokenRepo.ListTokens()
	if err != nil {
		return err
	}
	if len(tokens) == 0 {
		return nil // Nada para monitorar
	}

	// 2. Prepara mapa de decimais e lista de endereços
	decimalsMap := make(map[string]int)
	var addresses []string

	for _, t := range tokens {
		addrLower := strings.ToLower(t.Address)
		addresses = append(addresses, t.Address) // Endereço original (com Checksum ou não)
		decimalsMap[addrLower] = t.Decimals
	}

	// 3. Busca logs na Blockchain
	logs, err := s.ethClient.FetchTransfers(fromBlock, toBlock, addresses)
	if err != nil {
		return err
	}

	if len(logs) > 0 {
		slog.Info("Logs encontrados", "qtd", len(logs), "block_range", []int64{fromBlock, toBlock})
	}

	// 4. Processa cada log
	for _, vLog := range logs {
		if len(vLog.Topics) < 3 {
			continue // Log inválido (não é Transfer padrão)
		}

		contractAddr := strings.ToLower(vLog.Address.Hex())

		// Descobre os decimais deste contrato específico
		decimals := 18 // Default seguro
		if d, ok := decimalsMap[contractAddr]; ok {
			decimals = d
		}

		// Decodifica o valor (Core Logic)
		readableAmount := BytesToFloat(vLog.Data, decimals)

		// Monta objeto de domínio
		tx := &domain.Transfer{
			TxHash:          vLog.TxHash.Hex(),
			From:            common.HexToAddress(vLog.Topics[1].Hex()).Hex(),
			To:              common.HexToAddress(vLog.Topics[2].Hex()).Hex(),
			Amount:          readableAmount,
			BlockNumber:     vLog.BlockNumber,
			Timestamp:       time.Now(), // Idealmente pegariamos o timestamp do bloco, mas requer outra chamada RPC. Time.now serve por enquanto.
			ContractAddress: vLog.Address.Hex(),
		}

		// Salva
		if err := s.txRepo.SaveTransfer(tx); err != nil {
			slog.Error("Erro ao salvar tx", "hash", tx.TxHash, "err", err)
		} else {
			// Atualiza checkpoint individual do token
			s.txRepo.UpdateLastProcessedBlock(tx.ContractAddress, tx.BlockNumber)
		}
	}
	return nil
}
