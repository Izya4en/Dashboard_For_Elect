package grpc

import (
	"context"
	"techdash/internal/domain"
	"techdash/internal/transport/grpc/pb" // Путь к сгенерированным файлам

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type analyticsClient struct {
	client pb.SpatialAnalyticsClient
}

// NewAnalyticsClient подключается к порту Python (например, localhost:50051)
func NewAnalyticsClient(targetUrl string) (domain.SpatialAnalyticsClient, error) {
	conn, err := grpc.Dial(targetUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &analyticsClient{
		client: pb.NewSpatialAnalyticsClient(conn),
	}, nil
}

func (c *analyticsClient) GetHeatmap(ctx context.Context, points []domain.GeoPoint) ([]domain.GeoPoint, error) {
	// 1. Конвертируем доменные точки Go в формат protobuf
	var reqPoints []*pb.Point
	for _, p := range points {
		reqPoints = append(reqPoints, &pb.Point{Lat: p.Lat, Lon: p.Lon})
	}

	req := &pb.HeatmapRequest{Points: reqPoints}

	// 2. Отправляем в PYTHON!
	resp, err := c.client.CalculateHeatmap(ctx, req)
	if err != nil {
		return nil, err
	}

	// 3. Конвертируем ответ Python обратно в доменные точки
	var result []domain.GeoPoint
	for _, p := range resp.DensityPoints {
		result = append(result, domain.GeoPoint{
			Lat:    p.Lat,
			Lon:    p.Lon,
			Weight: p.Weight,
		})
	}

	return result, nil
}
