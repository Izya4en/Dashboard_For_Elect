package http

import (
	"techdash/internal/domain"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	orgUsecase domain.OrganizationUsecase,
	srvUsecase domain.ServiceTypeUsecase,
	actUsecase domain.ActUsecase,
	dashboardUsecase domain.DashboardUsecase,
	reportUsecase domain.ReportUsecase,
	analyticsUsecase domain.AnalyticsUsecase,
) {
	NewOrganizationHandler(r, orgUsecase)
	NewServiceTypeHandler(r, srvUsecase)
	NewActHandler(r, actUsecase)

	NewDashboardHandler(r, dashboardUsecase)
	NewReportHandler(r, reportUsecase)
	NewAnalyticsHandler(r, analyticsUsecase)
}
