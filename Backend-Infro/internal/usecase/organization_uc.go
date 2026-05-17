package usecase

import (
	"errors"
	"techdash/internal/domain"
)

type organizationUsecase struct {
	repo domain.OrganizationRepository
}

func NewOrganizationUsecase(r domain.OrganizationRepository) domain.OrganizationUsecase {
	return &organizationUsecase{repo: r}
}

func (u *organizationUsecase) GetAll() ([]domain.Organization, error) {
	return u.repo.GetAll()
}

func (u *organizationUsecase) Create(org *domain.Organization) error {
	// Пример бизнес-правила: БИН должен быть 12 символов
	if len(org.BIN) != 12 {
		return errors.New("БИН должен содержать ровно 12 символов")
	}
	return u.repo.Create(org)
}

func (u *organizationUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
