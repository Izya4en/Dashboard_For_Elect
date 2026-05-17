import grpc
from concurrent import futures
import numpy as np
from scipy.stats import gaussian_kde

# Импортируем сгенерированные файлы
import analytics_pb2
import analytics_pb2_grpc

class SpatialAnalyticsServicer(analytics_pb2_grpc.SpatialAnalyticsServicer):
    def CalculateHeatmap(self, request, context):
        print(f"Получено точек для анализа: {len(request.points)}")
        
        # Если точек нет или их слишком мало, просто возвращаем как есть с весом 1
        if len(request.points) < 3:
            out_points = [
                analytics_pb2.Point(lat=p.lat, lon=p.lon, weight=1.0) 
                for p in request.points
            ]
            return analytics_pb2.HeatmapResponse(density_points=out_points)

        # Вытаскиваем координаты в numpy массивы
        lats = np.array([p.lat for p in request.points])
        lons = np.array([p.lon for p in request.points])

        # Выполняем KDE (Ядерную оценку плотности)
        # Это определит "густоту" точек в пространстве
        positions = np.vstack([lats, lons])
        kde = gaussian_kde(positions)
        
        # Считаем плотность для каждой точки
        densities = kde(positions)

        # Нормализуем значения от 0 до 1 (чтобы фронтенду было удобно красить карту)
        max_density = max(densities)
        if max_density > 0:
            densities = densities / max_density

        # Формируем ответ
        out_points = []
        for lat, lon, density in zip(lats, lons, densities):
            out_points.append(analytics_pb2.Point(lat=lat, lon=lon, weight=float(density)))

        print("Анализ завершен, данные отправлены обратно.")
        return analytics_pb2.HeatmapResponse(density_points=out_points)

def serve():
    # Создаем многопоточный сервер
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    analytics_pb2_grpc.add_SpatialAnalyticsServicer_to_server(SpatialAnalyticsServicer(), server)
    
    # Слушаем порт 50051 (стандартный для gRPC)
    server.add_insecure_port('[::]:50051')
    print("Python Analytics Service запущен на порту 50051...")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()