package http

import (
	"net/http"
	"strconv"
	"techdash/internal/domain"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	usecase domain.OrganizationUsecase
}

func NewOrganizationHandler(r *gin.Engine, us domain.OrganizationUsecase) {
	handler := &OrganizationHandler{
		usecase: us,
	}

	// Группируем роуты для API v1
	v1 := r.Group("/api/v1/organizations")
	{
		v1.GET("", handler.GetAll)
		v1.POST("", handler.Create)
		v1.DELETE(":id", handler.Delete)
	}
}

func (h *OrganizationHandler) GetAll(c *gin.Context) {
	orgs, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orgs)
}

func (h *OrganizationHandler) Create(c *gin.Context) {
	var org domain.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if err := h.usecase.Create(&org); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, org)
}

func (h *OrganizationHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор организации"})
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
