package http

import (
	"net/http"
	"strconv"
	"strings"

	"go-chain/internal/core/service/admin"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *admin.AdminService
}

func NewHandler(s *admin.AdminService) *Handler {
	return &Handler{service: s}
}

// --- ROTAS DE CONTRATOS ---

// ListContracts godoc
// @Summary Lista todos os tokens monitorados
// @Router /api/v1/contracts [get]
func (h *Handler) ListContracts(c *gin.Context) {
	tokens, err := h.service.ListTokens()
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Erro ao listar contratos")
		return
	}
	SendSuccess(c, tokens)
}

type AddTokenReq struct {
	Address  string `json:"address" binding:"required"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol" binding:"required"`
	Decimals int    `json:"decimals"` // Opcional, default no service poderia ser tratado
}

// AddContract godoc
// @Summary Adiciona um novo token para monitoramento
// @Router /api/v1/contracts [post]
func (h *Handler) AddContract(c *gin.Context) {
	var req AddTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, http.StatusBadRequest, "Dados inválidos: "+err.Error())
		return
	}

	// Validação básica de decimals
	if req.Decimals <= 0 {
		req.Decimals = 18 // Default se não vier
	}

	err := h.service.AddToken(req.Address, req.Name, req.Symbol, req.Decimals)
	if err != nil {
		// Verifica se é erro de duplicidade (simplificado)
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			SendError(c, http.StatusConflict, "Token já cadastrado")
			return
		}
		SendError(c, http.StatusInternalServerError, "Erro ao salvar token: "+err.Error())
		return
	}

	SendSuccess(c, "Token adicionado com sucesso")
}

// RemoveContract godoc
// @Summary Para de monitorar um token
// @Router /api/v1/contracts/:address [delete]
func (h *Handler) RemoveContract(c *gin.Context) {
	addr := c.Param("address")
	if addr == "" {
		SendError(c, http.StatusBadRequest, "Endereço obrigatório")
		return
	}

	if err := h.service.RemoveToken(addr); err != nil {
		SendError(c, http.StatusInternalServerError, "Erro ao remover token")
		return
	}

	SendSuccess(c, "Token removido")
}

// --- ROTAS DO DASHBOARD ---

// GetTransfers godoc
// @Summary Busca transferências paginadas de um contrato específico
// @Router /api/v1/transfers [get]
func (h *Handler) GetTransfers(c *gin.Context) {
	// Query params
	contract := c.Query("contract")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if contract == "" {
		SendError(c, http.StatusBadRequest, "Parâmetro 'contract' é obrigatório")
		return
	}

	// Como o GetDashboardData retorna (transfers, stats, error), e aqui queremos só transfers
	// Precisamos adaptar ou chamar o repositório direto?
	// Pela arquitetura, o handler fala com o Service.
	// Vamos usar o GetDashboardData e ignorar o stats aqui, ou criar um metodo só pra transfers no service.
	// Para simplificar e seguir o AdminService que criamos:

	transfers, _, err := h.service.GetDashboardData(contract, page, limit)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Erro ao buscar transferências")
		return
	}

	SendSuccess(c, transfers)
}

// GetStats godoc
// @Summary Busca estatísticas (Volume, Whale, Count)
// @Router /api/v1/stats [get]
func (h *Handler) GetStats(c *gin.Context) {
	contract := c.Query("contract")
	if contract == "" {
		SendError(c, http.StatusBadRequest, "Parâmetro 'contract' é obrigatório")
		return
	}

	// Reutilizamos a lógica do service. Paginação 1, limit 1 só pra pegar o stats rápido.
	_, stats, err := h.service.GetDashboardData(contract, 1, 1)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Erro ao calcular estatísticas")
		return
	}

	SendSuccess(c, stats)
}
