package domain

// DashboardSummary - данные для главного экрана
type DashboardSummary struct {
	TotalOrganizations int   `json:"total_organizations"`
	TotalActs          int   `json:"total_acts"`
	TotalServices      int   `json:"total_services"`
	RecentActs         []Act `json:"recent_acts"`
}

type DashboardUsecase interface {
	GetSummary() (*DashboardSummary, error)
}
