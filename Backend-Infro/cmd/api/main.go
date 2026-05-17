package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Импорт драйвера Postgres

	"techdash/internal/repository/postgres"
	mygrpc "techdash/internal/transport/grpc" // ---> НОВОЕ: Импорт нашего gRPC пакета
	transport "techdash/internal/transport/http"
	"techdash/internal/usecase"
)

func main() {
	// 1. Подключение к БД
	connStr := "host=localhost port=5432 user=postgres password=root dbname=techdash sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("БД недоступна: %v", err)
	}

	if err := initDB(db); err != nil {
		log.Fatalf("Не удалось инициализировать БД: %v", err)
	}

	// 2. Инициализация слоев (Dependency Injection)

	// --- Слой данных (Repository) ---
	orgRepo := postgres.NewOrganizationRepository(db)
	actRepo := postgres.NewActRepository(db)
	srvRepo := postgres.NewServiceTypeRepository(db)

	// ---> НОВОЕ: Слой gRPC Клиентов <---
	// Подключаемся к Python микросервису
	grpcClient, err := mygrpc.NewAnalyticsClient("localhost:50051")
	if err != nil {
		log.Fatalf("Не удалось подключиться к Python Analytics: %v", err)
	}

	// --- Слой бизнес-логики (Usecase) ---
	orgUsecase := usecase.NewOrganizationUsecase(orgRepo)
	actUsecase := usecase.NewActUsecase(actRepo)
	srvUsecase := usecase.NewServiceTypeUsecase(srvRepo)

	// Дашборд собирает данные из трех репозиториев
	dashboardUsecase := usecase.NewDashboardUsecase(orgRepo, actRepo, srvRepo)

	// Сервис отчетов использует данные об актах для генерации Excel
	reportUsecase := usecase.NewReportUsecase(actRepo)

	// ---> НОВОЕ: Инициализация Usecase для аналитики
	analyticsUsecase := usecase.NewAnalyticsUsecase(grpcClient) // Если в будущем добавите orgRepo для реальных координат, передайте его сюда вторым параметром

	// 3. Настройка роутера
	router := gin.Default()

	// Настраиваем CORS (разрешаем запросы с фронтенда)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 4. Регистрация хендлеров (Transport Layer)
	transport.RegisterRoutes(router, orgUsecase, srvUsecase, actUsecase, dashboardUsecase, reportUsecase, analyticsUsecase)

	// 5. Запуск сервера
	log.Println("Бэкенд techDash запущен на порту 8080...")
	log.Println("Доступные эндпоинты:")
	log.Println(" - GET  /api/v1/dashboard      (Сводка для главной)")
	log.Println(" - GET  /api/v1/reports/export (Скачать Excel отчет)")
	log.Println(" - GET  /api/v1/analytics/heatmap (Тепловая карта плотности)") // ---> НОВОЕ

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func initDB(db *sql.DB) error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}

	sqlPath := filepath.Join(root, "init.sql")
	data, err := os.ReadFile(sqlPath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(data))
	return err
}
