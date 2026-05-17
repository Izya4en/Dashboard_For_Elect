package domain

// Organization - сущность клиента
type Organization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	BIN  string `json:"bin"`
}

// OrganizationRepository описывает работу с базой данных
type OrganizationRepository interface {
	GetAll() ([]Organization, error)
	Create(org *Organization) error
	Delete(id int) error
	GetCount() (int, error)
}

// OrganizationUsecase описывает бизнес-логику
type OrganizationUsecase interface {
	GetAll() ([]Organization, error)
	Create(org *Organization) error
	Delete(id int) error
}
