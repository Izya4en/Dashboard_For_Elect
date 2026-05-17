package domain

// Act - структура для работы с транзакциями
type Act struct {
	ID               int    `json:"id"`
	TestDate         string `json:"date"`            // Дата в формате YYYY-MM-DD
	OrganizationID   int    `json:"organization_id"` // Для сохранения
	OrganizationName string `json:"organization"`    // Для вывода в таблице
	OrganizationBIN  string `json:"bin"`             // Для вывода в таблице
	ServiceTypeID    int    `json:"service_type_id"` // Для сохранения
	ServiceName      string `json:"service"`         // Для вывода в таблице
}

type ActRepository interface {
	GetAll() ([]Act, error)
	Create(act *Act) error
	GetCount() (int, error)
	GetRecent(limit int) ([]Act, error)
}

type ActUsecase interface {
	GetAll() ([]Act, error)
	Create(act *Act) error
}
