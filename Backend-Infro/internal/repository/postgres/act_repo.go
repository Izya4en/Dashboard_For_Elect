package postgres

import (
	"database/sql"
	"techdash/internal/domain"
)

type actRepo struct {
	db *sql.DB
}

func NewActRepository(db *sql.DB) domain.ActRepository {
	return &actRepo{db: db}
}

func (r *actRepo) GetAll() ([]domain.Act, error) {
	// SQL-запрос, который собирает данные из трех таблиц сразу
	query := `
		SELECT 
			a.id, a.test_date, 
			o.name, o.bin, 
			s.name
		FROM acts a
		JOIN organizations o ON a.organization_id = o.id
		JOIN service_types s ON a.service_type_id = s.id
		ORDER BY a.test_date DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var acts []domain.Act
	for rows.Next() {
		var a domain.Act
		// Сканируем дату как строку (Postgres отдаст ее в формате RFC3339)
		if err := rows.Scan(&a.ID, &a.TestDate, &a.OrganizationName, &a.OrganizationBIN, &a.ServiceName); err != nil {
			return nil, err
		}
		// Обрезаем время, оставляем только дату YYYY-MM-DD
		a.TestDate = a.TestDate[:10]
		acts = append(acts, a)
	}
	return acts, nil
}

func (r *actRepo) Create(act *domain.Act) error {
	query := "INSERT INTO acts (test_date, organization_id, service_type_id) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, act.TestDate, act.OrganizationID, act.ServiceTypeID).Scan(&act.ID)
}

func (r *actRepo) GetCount() (int, error) {
	var count sql.NullInt64 // Безопасный тип
	err := r.db.QueryRow("SELECT COUNT(*) FROM acts").Scan(&count)

	if err != nil {
		return 0, nil // Игнорируем ошибку и отдаем 0
	}
	if count.Valid {
		return int(count.Int64), nil // Если число есть, отдаем его
	}
	return 0, nil
}

func (r *actRepo) GetRecent(limit int) ([]domain.Act, error) {
	query := `
		SELECT a.id, a.test_date, o.name, o.bin, s.name
		FROM acts a
		JOIN organizations o ON a.organization_id = o.id
		JOIN service_types s ON a.service_type_id = s.id
		ORDER BY a.test_date DESC LIMIT $1
	`
	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var acts []domain.Act
	for rows.Next() {
		var a domain.Act
		if err := rows.Scan(&a.ID, &a.TestDate, &a.OrganizationName, &a.OrganizationBIN, &a.ServiceName); err != nil {
			return nil, err
		}
		a.TestDate = a.TestDate[:10]
		acts = append(acts, a)
	}
	return acts, nil
}
