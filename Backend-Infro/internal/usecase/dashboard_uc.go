package usecase

import "techdash/internal/domain"

type dashboardUsecase struct {
	orgRepo domain.OrganizationRepository
	actRepo domain.ActRepository
	srvRepo domain.ServiceTypeRepository // <--- Добавили репозиторий услуг
}

// В конструктор теперь передаем ТРИ репозитория
func NewDashboardUsecase(or domain.OrganizationRepository, ar domain.ActRepository, sr domain.ServiceTypeRepository) domain.DashboardUsecase {
	return &dashboardUsecase{
		orgRepo: or,
		actRepo: ar,
		srvRepo: sr, // <--- Сохраняем
	}
}

func (u *dashboardUsecase) GetSummary() (*domain.DashboardSummary, error) {
	orgCount, err := u.orgRepo.GetCount()
	if err != nil {
		return nil, err
	}

	actCount, err := u.actRepo.GetCount()
	if err != nil {
		return nil, err
	}

	// ---> НОВОЕ: Считаем реальное количество услуг из БД
	srvCount, err := u.srvRepo.GetCount()
	if err != nil {
		return nil, err
	}

	recentActs, err := u.actRepo.GetRecent(5)
	if err != nil {
		return nil, err
	}

	return &domain.DashboardSummary{
		TotalOrganizations: orgCount,
		TotalActs:          actCount,
		TotalServices:      srvCount, // <--- Ставим переменную вместо цифры 6
		RecentActs:         recentActs,
	}, nil
}
