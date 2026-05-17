package usecase

import (
	"context"
	"techdash/internal/domain"
)

type analyticsUsecase struct {
	grpcClient domain.SpatialAnalyticsClient
	// Сюда можно добавить orgRepo domain.OrganizationRepository, чтобы брать реальные координаты
}

func NewAnalyticsUsecase(client domain.SpatialAnalyticsClient) domain.AnalyticsUsecase {
	return &analyticsUsecase{grpcClient: client}
}

func (u *analyticsUsecase) GenerateHeatmap(ctx context.Context) ([]domain.GeoPoint, error) {
	// Для примера создадим пару тестовых координат в Астане
	// В реальности здесь будет вызов базы данных: orgs, err := u.orgRepo.GetAll()
	points := []domain.GeoPoint{
		{Lat: 51.169392, Lon: 71.449074},
		{Lat: 51.170000, Lon: 71.450000},
		{Lat: 51.168000, Lon: 71.448000},
	}

	// Отправляем в Python на расчет плотности
	return u.grpcClient.GetHeatmap(ctx, points)
}
