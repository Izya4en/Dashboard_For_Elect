package domain

import "context"

// GeoPoint - точка для тепловой карты
type GeoPoint struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Weight float64 `json:"weight"` // Вес, который рассчитает Python
}

// SpatialAnalyticsClient - интерфейс для связи с Python
type SpatialAnalyticsClient interface {
	GetHeatmap(ctx context.Context, points []GeoPoint) ([]GeoPoint, error)
}

// AnalyticsUsecase - бизнес-логика
type AnalyticsUsecase interface {
	GenerateHeatmap(ctx context.Context) ([]GeoPoint, error)
}
