package courier

import (
	"github.com/AntonyIS/go-foods/internal/core/domain"
	"github.com/AntonyIS/go-foods/internal/core/ports"
)

type courierService struct {
	repo ports.CourierRepository
}

func NewCourierService(repo ports.CourierRepository) ports.CourierService {
	return &courierService{
		repo,
	}
}

func (cs *courierService) CreateCourier(courier *domain.Courier) (*domain.Courier, error) {
	return cs.repo.CreateCourier(courier)
}

func (cs *courierService) GetCouriers() (*[]domain.Courier, error) {
	return cs.repo.GetCouriers()
}

func (cs *courierService) GetCourier(id string) (*domain.Courier, error) {
	return cs.repo.GetCourier(id)
}

func (cs *courierService) UpdateCourier(courier *domain.Courier) (*domain.Courier, error) {
	return cs.repo.UpdateCourier(courier)
}

func (cs *courierService) DeleteCourier(id string) error {
	return cs.repo.DeleteCourier(id)
}
