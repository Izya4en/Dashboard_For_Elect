package postgres

import (
	"database/sql"
	"techdash/internal/domain"
)

type serviceTypeRepo struct {
	db *sql.DB
}

func NewServiceTypeRepository(db *sql.DB) domain.ServiceTypeRepository {
	return &serviceTypeRepo{db: db}
}

func (r *serviceTypeRepo) GetAll() ([]domain.ServiceType, error) {
	rows, err := r.db.Query("SELECT id, name FROM service_types ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []domain.ServiceType
	for rows.Next() {
		var st domain.ServiceType
		if err := rows.Scan(&st.ID, &st.Name); err != nil {
			return nil, err
		}
		types = append(types, st)
	}
	return types, nil
}

func (r *serviceTypeRepo) Create(st *domain.ServiceType) error {
	query := `INSERT INTO service_types (name) VALUES ($1) ON CONFLICT (name) DO NOTHING RETURNING id`
	err := r.db.QueryRow(query, st.Name).Scan(&st.ID)
	if err == sql.ErrNoRows {
		return r.db.QueryRow("SELECT id FROM service_types WHERE name = $1", st.Name).Scan(&st.ID)
	}
	return err
}

func (r *serviceTypeRepo) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM service_types WHERE id = $1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *serviceTypeRepo) GetCount() (int, error) {
	var count sql.NullInt64
	err := r.db.QueryRow("SELECT COUNT(*) FROM service_types").Scan(&count)

	if err != nil {
		return 0, nil
	}
	if count.Valid {
		return int(count.Int64), nil
	}
	return 0, nil
}
