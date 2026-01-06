package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	client *ethclient.Client
}

func NewEthClient(rpcURL string) (*EthClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &EthClient{client: client}, nil
}

// GetLatestBlock retorna o número do bloco atual
func (c *EthClient) GetLatestBlock() (uint64, error) {
	return c.client.BlockNumber(context.Background())
}

// FetchTransfers busca logs brutos (Event Transfer) de uma lista de endereços
func (c *EthClient) FetchTransfers(fromBlock, toBlock int64, addresses []string) ([]types.Log, error) {
	// Assinatura do evento Transfer: ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
	transferSig := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

	// Converte strings para common.Address
	var addrList []common.Address
	for _, a := range addresses {
		addrList = append(addrList, common.HexToAddress(a))
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: addrList,
		Topics:    [][]common.Hash{{transferSig}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.client.FilterLogs(ctx, query)
}
