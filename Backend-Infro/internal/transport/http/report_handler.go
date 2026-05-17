package http

import (
	"io"
	"net/http"
	"techdash/internal/domain"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	usecase domain.ReportUsecase
}

func NewReportHandler(r *gin.Engine, us domain.ReportUsecase) {
	handler := &ReportHandler{usecase: us}

	v1 := r.Group("/api/v1/reports")
	{
		v1.GET("/export", handler.ExportExcel)
	}
}

func (h *ReportHandler) ExportExcel(c *gin.Context) {
	fileReader, fileName, err := h.usecase.GenerateActsExcel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем заголовки, чтобы браузер понял, что это файл для скачивания
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	// Копируем данные из буфера в ответ HTTP
	io.Copy(c.Writer, fileReader)
}
