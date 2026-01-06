package config

import (
	"log/slog"

	"go-chain/internal/adapter/eth"
	"go-chain/internal/adapter/http"
	"go-chain/internal/adapter/storage/postgres"
	"go-chain/internal/core/service/admin"
	"go-chain/internal/core/service/monitor"
)

// Container segura todas as dependências da aplicação
type Container struct {
	MonitorService *monitor.MonitorService
	HTTPHandler    *http.Handler
	DB             *postgres.PostgresRepository
}

// NewContainer inicializa todas as dependências
func NewContainer(cfg *Config) *Container {
	// 1. Adapters (Banco e Blockchain)
	dbRepo, err := postgres.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		slog.Error("Falha ao conectar no banco", "err", err)
		panic(err)
	}

	ethClient, err := eth.NewEthClient(cfg.RPCURL)
	if err != nil {
		slog.Error("Falha ao conectar no RPC Ethereum", "err", err)
		panic(err)
	}

	// 2. Services (Core)
	// O PostgresRepo implementa tanto TokenRepo quanto TransferRepo
	monitorSvc := monitor.NewMonitorService(ethClient, dbRepo, dbRepo)
	adminSvc := admin.NewAdminService(dbRepo, dbRepo)

	// 3. Handlers (HTTP)
	httpHandler := http.NewHandler(adminSvc)

	return &Container{
		MonitorService: monitorSvc,
		HTTPHandler:    httpHandler,
		DB:             dbRepo,
	}
}
