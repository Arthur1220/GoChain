package monitor

import (
	"log/slog"
	"time"
)

// Start inicia o loop infinito de monitoramento
func (s *MonitorService) Start() {
	slog.Info("🚀 Monitor Worker Iniciado")

	// Pega bloco inicial (pode ser hardcoded ou do banco, aqui simplifiquei para pegar o atual da rede na partida)
	// Melhoria futura: Pegar o MENOR 'last_block' de todos os tokens no banco.
	currentBlock, err := s.ethClient.GetLatestBlock()
	if err != nil {
		slog.Error("Falha crítica ao iniciar worker", "err", err)
		return
	}
	slog.Info("Começando sincronização a partir do bloco", "block", currentBlock)

	// Loop Infinito
	for {
		// 1. Qual o topo da cadeia agora?
		latestBlock, err := s.ethClient.GetLatestBlock()
		if err != nil {
			slog.Error("Erro RPC", "err", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// 2. Estamos sincronizados?
		if uint64(currentBlock) > latestBlock {
			time.Sleep(6 * time.Second) // Espera novo bloco (Ethereum block time ~12s)
			continue
		}

		// 3. Define tamanho do lote (Batch Size)
		// Processamos no máximo 10 blocos por vez para não estourar RPC
		endBlock := uint64(currentBlock) + 10
		if endBlock > latestBlock {
			endBlock = latestBlock
		}

		// 4. Processa
		err = s.ProcessRange(int64(currentBlock), int64(endBlock))
		if err != nil {
			slog.Error("Falha ao processar range", "err", err)
			time.Sleep(5 * time.Second) // Backoff em caso de erro
			continue
		}

		// 5. Avança cursor
		currentBlock = endBlock + 1
	}
}
