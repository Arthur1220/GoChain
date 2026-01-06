package http

import "github.com/gin-gonic/gin"

// Response é o formato padrão de resposta JSON
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SendSuccess envia uma resposta 200 OK com dados
func SendSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Success: true,
		Data:    data,
	})
}

// SendError envia uma resposta de erro (400 ou 500)
func SendError(c *gin.Context, status int, msg string) {
	c.JSON(status, Response{
		Success: false,
		Error:   msg,
	})
}
