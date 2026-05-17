package http

import (
	"net/http"
	"techdash/internal/domain"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	usecase domain.DashboardUsecase
}

func NewDashboardHandler(r *gin.Engine, us domain.DashboardUsecase) {
	handler := &DashboardHandler{usecase: us}

	// Регистрируем роут
	r.GET("/api/v1/dashboard", handler.GetSummary)
}

func (h *DashboardHandler) GetSummary(c *gin.Context) {
	summary, err := h.usecase.GetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}
