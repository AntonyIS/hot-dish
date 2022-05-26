package customer

import (
	"github.com/AntonyIS/go-foods/internal/core/domain"
	"github.com/AntonyIS/go-foods/internal/core/ports"
)

type customerService struct {
	repo ports.CustomerRepository
}

func NewCustomerService(repo ports.CustomerRepository) ports.CustomerService {
	return &customerService{
		repo,
	}
}

func (cs customerService) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	return cs.repo.CreateCustomer(customer)
}

func (cs customerService) GetCustomers() (*[]domain.Customer, error) {
	return cs.repo.GetCustomers()
}

func (cs customerService) GetCustomer(id string) (*domain.Customer, error) {
	return cs.repo.GetCustomer(id)
}

func (cs customerService) UpdateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	return cs.repo.UpdateCustomer(customer)
}

func (cs customerService) DeleteCustomer(id string) error {
	return cs.repo.DeleteCustomer(id)
}
