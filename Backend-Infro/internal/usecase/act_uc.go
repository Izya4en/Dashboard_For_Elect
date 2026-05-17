package usecase

import "techdash/internal/domain"

type actUsecase struct {
	repo domain.ActRepository
}

func NewActUsecase(r domain.ActRepository) domain.ActUsecase {
	return &actUsecase{repo: r}
}

func (u *actUsecase) GetAll() ([]domain.Act, error) {
	return u.repo.GetAll()
}

func (u *actUsecase) Create(act *domain.Act) error {
	// Здесь можно добавить проверки, например, не из будущего ли дата
	return u.repo.Create(act)
}
