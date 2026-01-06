package ports

import "github.com/ethereum/go-ethereum/core/types"

// BlockchainService define como interagimos com a rede (Ethereum)
type BlockchainService interface {
	// Retorna o número do bloco mais recente da rede
	GetLatestBlock() (uint64, error)

	// Busca logs (eventos) de múltiplos endereços em um range de blocos
	FetchTransfers(fromBlock, toBlock int64, addresses []string) ([]types.Log, error)
}
