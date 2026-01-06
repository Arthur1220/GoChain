package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter inicializa o Gin e define as rotas
func InitRouter(handler *Handler) *gin.Engine {
	// Define modo Release se necessário (ou debug por padrão)
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// --- Configuração de CORS (Crucial para o Vue funcionar) ---
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Permite localhost:5173 e qualquer outro
	config.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// --- Rotas da API ---
	api := r.Group("/api/v1")
	{
		// Contratos
		api.GET("/contracts", handler.ListContracts)
		api.POST("/contracts", handler.AddContract)
		api.DELETE("/contracts/:address", handler.RemoveContract)

		// Dashboard
		api.GET("/transfers", handler.GetTransfers)
		api.GET("/stats", handler.GetStats)
	}

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "system": "Go-Chain Turbo"})
	})

	return r
}
