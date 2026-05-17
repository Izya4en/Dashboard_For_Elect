package usecase

import "techdash/internal/domain"

type serviceTypeUsecase struct {
	repo domain.ServiceTypeRepository
}

func NewServiceTypeUsecase(r domain.ServiceTypeRepository) domain.ServiceTypeUsecase {
	return &serviceTypeUsecase{repo: r}
}

func (u *serviceTypeUsecase) GetAll() ([]domain.ServiceType, error) {
	return u.repo.GetAll()
}

func (u *serviceTypeUsecase) Create(st *domain.ServiceType) error {
	return u.repo.Create(st)
}

func (u *serviceTypeUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
