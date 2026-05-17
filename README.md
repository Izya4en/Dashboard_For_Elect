# techDash — Дашборд для учёта электрических испытаний

Веб-приложение для управления актами технических испытаний и проверок электрооборудования. Позволяет вести реестр организаций, справочник видов услуг, регистрировать акты испытаний и анализировать статистику через интерактивные графики.

Система состоит из трёх независимых частей:
- **Front** — Vue 3 SPA (браузерный интерфейс)
- **Backend-Infro** — Go REST API (основной бэкенд, порт `8080`)
- **Backend-Analytic** — Python gRPC-микросервис (вычисления, порт `50051`)

---

## Содержание

- [Стек технологий](#стек-технологий)
- [Возможности системы](#возможности-системы)
- [Структура проекта](#структура-проекта)
- [Архитектура системы](#архитектура-системы)
- [Backend-Infro — Go REST API](#backend-infro--go-rest-api)
- [Backend-Analytic — Python gRPC](#backend-analytic--python-grpc)
- [База данных](#база-данных)
- [API-эндпоинты](#api-эндпоинты)
- [Установка и запуск](#установка-и-запуск)
- [Описание экранов](#описание-экранов)

---

## Стек технологий

### Фронтенд

| Слой         | Технология                                      |
|--------------|-------------------------------------------------|
| Фреймворк    | [Vue 3](https://vuejs.org/) (Composition API + `<script setup>`) |
| Сборщик      | [Vite 8](https://vitejs.dev/)                   |
| Стили        | [Tailwind CSS 3](https://tailwindcss.com/)      |
| Графики      | [ApexCharts](https://apexcharts.com/) + vue3-apexcharts |
| Состояние    | Реактивные `ref` Vue (без Pinia/Vuex)           |
| HTTP-клиент  | Нативный `fetch` API                            |

### Бэкенд

| Слой               | Технология                                    |
|--------------------|-----------------------------------------------|
| Go REST API        | [Gin](https://gin-gonic.com/) v1.12           |
| Драйвер PostgreSQL | `lib/pq`                                      |
| Excel-отчёты       | `excelize/v2`                                 |
| gRPC-клиент (Go)   | `google.golang.org/grpc` v1.81                |
| Python gRPC-сервер | `grpcio` + `grpcio-tools`                     |
| Математика         | `numpy` + `scipy` (gaussian_kde)              |
| База данных        | PostgreSQL 15 + PostGIS (Docker)              |
| Протокол связи     | Protocol Buffers (proto3)                     |

---

## Возможности системы

### Дашборд (главная страница)
- Сводные KPI-карточки: количество организаций, оказанных услуг, видов испытаний
- Таблица последних транзакций (акты испытаний) с поиском по организации, БИН и услуге
- Создание нового акта испытания через модальное окно
- Экспорт сводного отчёта в `.xlsx`

### Организации
- Полный список организаций с номерами БИН
- Добавление новой организации (наименование + 12-значный БИН)
- Редактирование и удаление существующих организаций
- Поиск по названию и БИН в реальном времени

### Справочник услуг
- Список видов испытаний и проверок
- Добавление, редактирование и удаление услуг
- Поиск по наименованию

### Аналитика
- KPI: всего транзакций, активных организаций, видов услуг, охватываемый период
- График динамики транзакций по месяцам (area chart)
- Горизонтальная диаграмма топ-12 услуг по частоте обращений (bar chart)
- Вертикальная диаграмма топ-15 организаций по числу транзакций (bar chart)
- Кнопка ручного обновления данных

---

## Структура проекта

```
Dashboard_For_Elect/
│
├── Front/                              # Фронтенд (Vue 3 + Vite)
│   ├── index.html
│   ├── package.json
│   ├── vite.config.js
│   └── src/
│       ├── main.js                    # Инициализация Vue-приложения
│       ├── App.vue                    # Корневой компонент: сайдбар, шапка, роутинг
│       ├── store.js                   # Глобальное состояние + все HTTP-вызовы к API
│       ├── mockData.js                # Тестовые данные для разработки без бэкенда
│       └── views/
│           ├── DashboardView.vue      # KPI-карточки + таблица транзакций
│           ├── OrganizationsView.vue  # CRUD организаций
│           ├── ServicesView.vue       # CRUD справочника услуг
│           └── AnalyticsView.vue      # Графики ApexCharts
│
├── Backend-Infro/                      # Go REST API (основной бэкенд)
│   ├── cmd/api/main.go                # Точка входа: DI, роутер, запуск сервера
│   ├── go.mod / go.sum
│   ├── init.sql                       # DDL + сидовые данные для БД
│   ├── docker-compose.yaml            # Поднимает PostgreSQL в Docker
│   └── internal/
│       ├── domain/                    # Интерфейсы и структуры (чистые типы)
│       │   ├── act.go
│       │   ├── organization.go
│       │   ├── service_type.go
│       │   ├── analytics.go           # GeoPoint, SpatialAnalyticsClient
│       │   ├── dashboard.go
│       │   └── report.go
│       ├── repository/postgres/       # Реализации интерфейсов (SQL-запросы)
│       ├── usecase/                   # Бизнес-логика
│       │   ├── act_uc.go
│       │   ├── organization_uc.go
│       │   ├── service_type_uc.go
│       │   ├── dashboard_uc.go
│       │   ├── report_uc.go
│       │   └── analytics_uc.go        # Вызывает Python через gRPC
│       └── transport/
│           ├── http/                  # Gin-хэндлеры + routes.go
│           └── grpc/                  # gRPC-клиент для связи с Python
│
└── Backend-Analytic/                   # Python gRPC-микросервис
    └── analytics/
        ├── analytics.proto            # Контракт (схема сообщений)
        ├── analytics_pb2.py           # Сгенерированные классы сообщений
        ├── analytics_pb2_grpc.py      # Сгенерированный gRPC-stub
        └── main.py                    # gRPC-сервер с KDE-вычислениями
```

---

## Архитектура системы

Система построена по принципу микросервисов. Три компонента общаются по разным протоколам:

```
 Браузер
    │  HTTP (REST JSON)
    ▼
┌─────────────────────────────────────────┐
│        Backend-Infro  (Go, :8080)       │
│                                         │
│  HTTP Handler → Usecase → Repository   │
│                                         │
│  ┌──────────┐        ┌───────────────┐  │
│  │ Gin      │        │  gRPC Client  │  │
│  │ Router   │        │  (аналитика)  │  │
│  └──────────┘        └───────┬───────┘  │
│         │                    │ gRPC     │
│         ▼                    ▼          │
│  ┌──────────────┐   ┌───────────────┐  │
│  │  PostgreSQL  │   │ Backend-Analytic│ │
│  │  (SQL, :5432)│   │ Python, :50051 │  │
│  └──────────────┘   └───────────────┘  │
└─────────────────────────────────────────┘
```

**Поток данных (фронтенд → бэкенд):**
1. `App.vue` при старте вызывает `fetchAllData()` из `store.js`
2. `store.js` параллельно делает 3 fetch-запроса: `/organizations`, `/services`, `/acts`
3. Ответы сохраняются в реактивные `ref`-переменные — все компоненты обновляются автоматически
4. Строка поиска хранится в `App.vue` и передаётся в каждый вью через prop `searchQuery`
5. При смене вкладки поиск очищается через `watch` на `activeTab`

**Поток данных (аналитика с тепловой картой):**
1. Go-бэкенд получает запрос `GET /api/v1/analytics/heatmap`
2. `analyticsUsecase` собирает координаты организаций из БД
3. Go отправляет список точек в Python через gRPC (метод `CalculateHeatmap`)
4. Python считает плотность распределения (KDE) и возвращает веса
5. Go возвращает обогащённый список точек фронтенду

---

## Backend-Infro — Go REST API

**Папка:** [Backend-Infro/](Backend-Infro/)  
**Фреймворк:** Gin  
**Порт:** `8080`

### Архитектура слоёв (Clean Architecture)

Go-бэкенд построен на принципе разделения ответственности: каждый слой знает только про своих соседей.

```
transport/http  →  usecase  →  repository/postgres  →  PostgreSQL
     ▲                │
     │           transport/grpc  →  Python Analytics
     │
  (Gin Router)
```

| Слой | Папка | Роль |
|------|-------|------|
| **Domain** | `internal/domain/` | Интерфейсы и структуры данных. Не зависит ни от чего. Здесь описывается «что умеет система» |
| **Repository** | `internal/repository/postgres/` | Реализация SQL-запросов к PostgreSQL. Знает о базе данных, но не знает о HTTP |
| **Usecase** | `internal/usecase/` | Бизнес-логика. Оркестрирует репозитории, вызывает gRPC. Не знает о HTTP |
| **Transport HTTP** | `internal/transport/http/` | Gin-хэндлеры. Парсит JSON, вызывает usecase, формирует ответ |
| **Transport gRPC** | `internal/transport/grpc/` | gRPC-клиент для вызова Python-сервиса |

### Как устроена точка входа (main.go)

```go
// 1. Подключение к PostgreSQL
db, _ := sql.Open("postgres", connStr)

// 2. Слой данных
orgRepo    := postgres.NewOrganizationRepository(db)
actRepo    := postgres.NewActRepository(db)
srvRepo    := postgres.NewServiceTypeRepository(db)

// 3. gRPC-клиент к Python
grpcClient, _ := mygrpc.NewAnalyticsClient("localhost:50051")

// 4. Бизнес-логика
orgUsecase       := usecase.NewOrganizationUsecase(orgRepo)
analyticsUsecase := usecase.NewAnalyticsUsecase(grpcClient)
// ...

// 5. Регистрация роутов в Gin
transport.RegisterRoutes(router, orgUsecase, ...)

// 6. Запуск
router.Run(":8080")
```

Такой подход называется **Dependency Injection (DI)** вручную: каждый слой получает свои зависимости через конструктор, а не создаёт их сам. Это упрощает тестирование и замену реализаций.

### Основные доменные модели

```go
// Акт испытания (транзакция)
type Act struct {
    ID               int    `json:"id"`
    TestDate         string `json:"date"`
    OrganizationID   int    `json:"organization_id"`
    OrganizationName string `json:"organization"`
    OrganizationBIN  string `json:"bin"`
    ServiceTypeID    int    `json:"service_type_id"`
    ServiceName      string `json:"service"`
}

// Организация (клиент)
type Organization struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    BIN  string `json:"bin"`
}

// Точка для тепловой карты
type GeoPoint struct {
    Lat    float64 `json:"lat"`
    Lon    float64 `json:"lon"`
    Weight float64 `json:"weight"` // заполняет Python
}
```

### Генерация Excel-отчёта

`reportUsecase` использует библиотеку `excelize/v2` для создания `.xlsx` файла прямо в памяти Go и отдаёт его как бинарный поток (`Content-Disposition: attachment`).

### CORS

В `main.go` настроен middleware, который разрешает запросы с любого origin (`*`), поддерживает методы `GET, POST, PUT, DELETE` и обрабатывает preflight `OPTIONS`-запросы. Это необходимо, так как фронтенд работает на другом порту (`5173`).

---

## Backend-Analytic — Python gRPC

**Папка:** [Backend-Analytic/analytics/](Backend-Analytic/analytics/)  
**Порт:** `50051` (gRPC)

### Что делает этот сервис

Принимает массив географических точек (широта + долгота) от Go-бэкенда и вычисляет для каждой точки **вес плотности** — насколько «густо» в этом месте расположены организации. Результат используется для построения тепловой карты.

### Алгоритм: Ядерная оценка плотности (KDE)

```
Входные данные:
  точка_1(lat=51.16, lon=71.44)
  точка_2(lat=51.18, lon=71.47)
  ...

                    ┌─────────────────────┐
                    │  gaussian_kde        │
                    │  (scipy.stats)       │
                    │                      │
                    │  Строит непрерывную  │
                    │  функцию плотности   │
                    │  по всем точкам      │
                    └──────────┬──────────┘
                               │
                               ▼
              density = kde(positions)  # для каждой точки
              density = density / max(density)  # нормализация [0, 1]

Выходные данные:
  точка_1(lat=51.16, lon=71.44, weight=0.93)  ← высокая плотность
  точка_2(lat=51.18, lon=71.47, weight=0.41)  ← средняя плотность
```

**KDE (Kernel Density Estimation)** — статистический метод, который для каждой точки считает, сколько других точек находится «рядом», учитывая расстояние. Чем больше соседей — тем выше вес. Gaussian KDE использует функцию нормального распределения как «ядро».

### Протокол взаимодействия (Protocol Buffers)

Файл [analytics.proto](Backend-Analytic/analytics/analytics.proto) — это контракт между Go и Python. Он описывает структуру сообщений и методы, не завися от языка.

```protobuf
service SpatialAnalytics {
  rpc CalculateHeatmap (HeatmapRequest) returns (HeatmapResponse);
}

message Point {
  double lat    = 1;  // широта
  double lon    = 2;  // долгота
  double weight = 3;  // вес (заполняет Python)
}

message HeatmapRequest  { repeated Point points        = 1; }
message HeatmapResponse { repeated Point density_points = 1; }
```

Из этого `.proto`-файла `protoc` автоматически генерирует:
- `analytics_pb2.py` — Python-классы для сообщений
- `analytics_pb2_grpc.py` — Python-stub для сервера

### Запуск Python-сервиса

```bash
cd Backend-Analytic/analytics

# Установить зависимости
pip install grpcio grpcio-tools numpy scipy

# Запустить сервер
python main.py
# → "Python Analytics Service запущен на порту 50051..."
```

### Особенности реализации

- **Многопоточность:** сервер использует `ThreadPoolExecutor(max_workers=10)` — может обрабатывать до 10 запросов параллельно
- **Защита от малых выборок:** если точек меньше 3, KDE вернёт ошибку (недостаточно данных для оценки плотности). Сервис обрабатывает этот случай и возвращает все точки с весом `1.0`
- **Нормализация:** итоговые веса делятся на максимальное значение, чтобы результат всегда был в диапазоне `[0, 1]` — удобно для окраски тепловой карты на фронте

---

## База данных

**СУБД:** PostgreSQL 15 + PostGIS (запускается через Docker Compose)  
**Файл схемы:** [Backend-Infro/init.sql](Backend-Infro/init.sql)

### Схема таблиц

```sql
organizations          service_types          acts
──────────────         ─────────────          ──────────────────
id   SERIAL PK         id   SERIAL PK         id              SERIAL PK
name VARCHAR(255)      name VARCHAR(255)       test_date       DATE
bin  VARCHAR(12)       (unique)                organization_id → organizations.id
lat  DOUBLE PRECISION                          service_type_id → service_types.id
lon  DOUBLE PRECISION
(lat/lon нужны для тепловой карты)
```

**Каскадное удаление:** если удалить организацию или вид услуги, все связанные акты удалятся автоматически (`ON DELETE CASCADE`).

### Запуск базы данных

```bash
cd Backend-Infro
docker-compose up -d
```

При первом запуске Docker автоматически выполнит `init.sql`, который:
1. Создаёт таблицы
2. Загружает 100 реалистичных организаций (ТОО, ИП из разных городов Казахстана)
3. Загружает 12 видов электротехнических испытаний
4. Загружает 100 тестовых актов за 2023 год

---

## API-эндпоинты

Базовый URL: `http://localhost:8080/api/v1`

| Метод    | Путь                      | Описание                              |
|----------|---------------------------|---------------------------------------|
| `GET`    | `/organizations`          | Список всех организаций               |
| `POST`   | `/organizations`          | Создать организацию                   |
| `DELETE` | `/organizations/:id`      | Удалить организацию по ID             |
| `GET`    | `/services`               | Список всех видов услуг               |
| `POST`   | `/services`               | Создать вид услуги                    |
| `DELETE` | `/services/:id`           | Удалить вид услуги по ID              |
| `GET`    | `/acts`                   | Список всех актов испытаний           |
| `POST`   | `/acts`                   | Создать акт испытания                 |
| `GET`    | `/reports/export`         | Скачать сводный отчёт (`.xlsx`)       |

### Схемы тел запросов (POST)

**Создать организацию:**
```json
{
  "name": "ООО ТехноЛаб",
  "bin": "060840000123"
}
```

**Создать вид услуги:**
```json
{
  "name": "Проверка электробезопасности"
}
```

**Создать акт испытания:**
```json
{
  "test_date": "2026-05-18",
  "organization_id": 1,
  "service_type_id": 2
}
```

---

## Установка и запуск

### Требования

| Инструмент | Версия      | Назначение          |
|------------|-------------|---------------------|
| Node.js    | >= 18       | Фронтенд            |
| Go         | >= 1.21     | REST API            |
| Python     | >= 3.9      | gRPC микросервис    |
| Docker     | любая       | PostgreSQL          |

### Шаг 1 — База данных (PostgreSQL)

```bash
cd Backend-Infro
docker-compose up -d
# Поднимает PostgreSQL на порту 5432
# Автоматически создаёт таблицы и загружает тестовые данные из init.sql
```

### Шаг 2 — Python Analytics Service

```bash
cd Backend-Analytic/analytics

pip install grpcio grpcio-tools numpy scipy

python main.py
# → Python Analytics Service запущен на порту 50051...
```

### Шаг 3 — Go REST API

```bash
cd Backend-Infro

go run ./cmd/api/main.go
# → Бэкенд techDash запущен на порту 8080...
```

> Если Go-сервер не может подключиться к Python — убедитесь, что шаг 2 выполнен первым.

### Шаг 4 — Фронтенд

```bash
cd Front

npm install
npm run dev
# → Открыть http://localhost:5173
```

### Доступные команды фронтенда

| Команда           | Описание                                      |
|-------------------|-----------------------------------------------|
| `npm run dev`     | Запустить дев-сервер с hot-reload             |
| `npm run build`   | Собрать продакшн-билд в папку `dist/`         |
| `npm run preview` | Предпросмотр продакшн-билда локально          |

### Смена адреса бэкенда

Если Go-бэкенд доступен не на `localhost:8080`, отредактируйте константу в [Front/src/store.js](Front/src/store.js):

```js
const API_BASE_URL = 'http://YOUR_HOST:PORT/api/v1'
```

### Порты по умолчанию

| Сервис              | Порт   |
|---------------------|--------|
| Фронтенд (Vite)     | `5173` |
| Go REST API         | `8080` |
| Python gRPC         | `50051`|
| PostgreSQL (Docker) | `5432` |

---

## Описание экранов

### Дашборд — Обзор испытаний
Главная страница системы. Показывает три KPI-карточки (организации, услуги, транзакции) и таблицу всех актов испытаний с колонками: ID, дата, организация, БИН, услуга. Кнопка **«Создать акт»** открывает модальное окно, где нужно выбрать дату, организацию и вид услуги. Кнопка **«Экспорт отчёта»** инициирует скачивание файла `report.xlsx` с бэкенда.

### Организации — Справочник клиентов
Таблица с полным реестром организаций. Каждая строка содержит ID, наименование и БИН. Доступны операции добавления (с валидацией 12-значного БИН), редактирования и удаления (с подтверждением). Поиск фильтрует по названию и БИН одновременно.

### Справочник услуг — Виды испытаний
Каталог всех видов технических услуг и испытаний. Простая таблица с кодом и наименованием. Поддерживает добавление, редактирование и удаление записей.

### Аналитика — Отчёты и графики
Агрегированная статистика по всей базе данных:
- **Area chart** — помесячная динамика количества актов
- **Horizontal bar** — топ-12 наиболее востребованных видов услуг
- **Vertical bar** — топ-15 наиболее активных организаций

Данные запрашиваются независимо от основного стора при монтировании компонента. Алгоритм определяет «доминирующий год» (год с наибольшим числом транзакций), чтобы исключить некорректные даты из временного графика.
