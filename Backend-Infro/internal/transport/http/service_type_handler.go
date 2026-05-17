package http

import (
	"net/http"
	"strconv"
	"techdash/internal/domain"

	"github.com/gin-gonic/gin"
)

type ServiceTypeHandler struct {
	usecase domain.ServiceTypeUsecase
}

func NewServiceTypeHandler(r *gin.Engine, us domain.ServiceTypeUsecase) {
	handler := &ServiceTypeHandler{usecase: us}

	v1 := r.Group("/api/v1/services")
	{
		v1.GET("", handler.GetAll)
		v1.POST("", handler.Create)
		v1.DELETE(":id", handler.Delete)
	}
}

func (h *ServiceTypeHandler) GetAll(c *gin.Context) {
	services, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

func (h *ServiceTypeHandler) Create(c *gin.Context) {
	var st domain.ServiceType
	if err := c.ShouldBindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := h.usecase.Create(&st); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, st)
}

func (h *ServiceTypeHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор услуги"})
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
