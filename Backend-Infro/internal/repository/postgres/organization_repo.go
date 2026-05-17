package postgres

import (
	"database/sql"
	"techdash/internal/domain"
)

type organizationRepo struct {
	db *sql.DB
}

func NewOrganizationRepository(db *sql.DB) domain.OrganizationRepository {
	return &organizationRepo{db: db}
}

func (r *organizationRepo) GetAll() ([]domain.Organization, error) {
	rows, err := r.db.Query("SELECT id, name, bin FROM organizations ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orgs []domain.Organization
	for rows.Next() {
		var o domain.Organization
		if err := rows.Scan(&o.ID, &o.Name, &o.BIN); err != nil {
			return nil, err
		}
		orgs = append(orgs, o)
	}
	return orgs, nil
}

func (r *organizationRepo) Create(org *domain.Organization) error {
	query := "INSERT INTO organizations (name, bin) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRow(query, org.Name, org.BIN).Scan(&org.ID)
}

func (r *organizationRepo) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM organizations WHERE id = $1", id)
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

func (r *organizationRepo) GetCount() (int, error) {
	var count sql.NullInt64
	err := r.db.QueryRow("SELECT COUNT(*) FROM organizations").Scan(&count)

	if err != nil {
		return 0, nil
	}
	if count.Valid {
		return int(count.Int64), nil
	}
	return 0, nil
}
