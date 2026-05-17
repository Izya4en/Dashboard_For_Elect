package domain

// ServiceType - вид испытания (услуга)
type ServiceType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ServiceTypeRepository interface {
	GetAll() ([]ServiceType, error)
	Create(st *ServiceType) error
	Delete(id int) error
	GetCount() (int, error) // <--- Для дашборда
}

type ServiceTypeUsecase interface {
	GetAll() ([]ServiceType, error)
	Create(st *ServiceType) error
	Delete(id int) error
}
