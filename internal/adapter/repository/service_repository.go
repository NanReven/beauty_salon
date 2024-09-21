package repository

import (
	"beauty_salon/internal/adapter/dto"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ServiceRepository struct {
	db *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (repo *ServiceRepository) GetAllServices() ([]dto.ServiceResponse, error) {
	var services []dto.ServiceResponse
	query := fmt.Sprintf("SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM %s JOIN %s ON (category_id = categories.id)", servicesTable, categoriesTable)
	if err := repo.db.Select(&services, query); err != nil {
		return nil, err
	}
	return services, nil
}

func (repo *ServiceRepository) GetServiceById(id int) (dto.ServiceResponse, error) {
	var services dto.ServiceResponse
	query := fmt.Sprintf("SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM %s JOIN %s ON (category_id = categories.id) WHERE services.id=$1", servicesTable, categoriesTable)
	if err := repo.db.Get(&services, query, id); err != nil {
		return services, err
	}
	return services, nil
}
