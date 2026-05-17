package http

import (
	"net/http"
	"techdash/internal/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type ActHandler struct {
	usecase domain.ActUsecase
}

func NewActHandler(r *gin.Engine, us domain.ActUsecase) {
	handler := &ActHandler{usecase: us}

	v1 := r.Group("/api/v1/acts")
	{
		v1.GET("", handler.GetAll)
		v1.POST("", handler.Create)
	}
}

func (h *ActHandler) GetAll(c *gin.Context) {
	acts, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acts)
}

type createActRequest struct {
	TestDate       string `json:"test_date"`
	OrganizationID int    `json:"organization_id"`
	ServiceTypeID  int    `json:"service_type_id"`
}

func (h *ActHandler) Create(c *gin.Context) {
	var req createActRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if req.TestDate == "" || req.OrganizationID == 0 || req.ServiceTypeID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Заполните все поля"})
		return
	}

	if _, err := time.Parse("2006-01-02", req.TestDate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат даты, ожидается YYYY-MM-DD"})
		return
	}

	act := domain.Act{
		TestDate:       req.TestDate,
		OrganizationID: req.OrganizationID,
		ServiceTypeID:  req.ServiceTypeID,
	}

	if err := h.usecase.Create(&act); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, act)
}
