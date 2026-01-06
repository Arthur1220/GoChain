package main

import (
	"log/slog"
	"os"

	"go-chain/config"
	"go-chain/internal/adapter/http"

	"github.com/joho/godotenv"
)

func main() {
	// Configuração de Logger (JSON estruturado é melhor para produção)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 1. Carrega variáveis de ambiente
	if err := godotenv.Load("backend/.env"); err != nil {
		// Tenta carregar da raiz se não achar na pasta backend
		_ = godotenv.Load(".env")
		slog.Warn("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	cfg := config.Load()

	// 2. Injeção de Dependência (Wiring)
	app := config.NewContainer(cfg)
	defer app.DB.Close() // Fecha banco ao encerrar

	// 3. Inicia o Worker de Monitoramento em Background (Goroutine)
	go func() {
		slog.Info("Iniciando Monitor Worker...")
		app.MonitorService.Start()
	}()

	// 4. Inicia o Servidor HTTP
	router := http.InitRouter(app.HTTPHandler)

	slog.Info("🚀 Servidor HTTP rodando", "port", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		slog.Error("Erro fatal no servidor HTTP", "err", err)
		os.Exit(1)
	}
}
