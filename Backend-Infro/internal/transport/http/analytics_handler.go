package http

import (
	"net/http"
	"techdash/internal/domain"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	usecase domain.AnalyticsUsecase
}

func NewAnalyticsHandler(r *gin.Engine, us domain.AnalyticsUsecase) {
	handler := &AnalyticsHandler{usecase: us}
	// Этот URL будет дергать ваш Nuxt фронтенд
	r.GET("/api/v1/analytics/heatmap", handler.GetHeatmap)
}

func (h *AnalyticsHandler) GetHeatmap(c *gin.Context) {
	points, err := h.usecase.GenerateHeatmap(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, points)
}
