package logger

import (
	"log/slog"
	"os"
)

// InitLogger configura o log para sair em formato JSON
func InitLogger() {
	// Cria um handler que escreve JSON no terminal (Stdout)
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Só mostra INFO, WARN e ERROR (ignora DEBUG)
	})

	logger := slog.New(jsonHandler)

	// Define esse logger como o padrão do sistema
	slog.SetDefault(logger)
}
