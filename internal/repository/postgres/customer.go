package postgres

import (
	"fmt"

	"github.com/AntonyIS/go-foods/internal/core/domain"
)

func (repo PostgresRepository) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {

	if err := repo.DB.Create(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
func (repo PostgresRepository) GetCustomers() (*[]domain.Customer, error) {

	customer := []domain.Customer{}
	if result := repo.DB.Find(&customer); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &customer, nil
}

func (repo PostgresRepository) GetCustomer(id string) (*domain.Customer, error) {
	customer := domain.Customer{}
	if result := repo.DB.Find(&customer); result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (repo PostgresRepository) UpdateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	if result := repo.DB.Save(customer); result.Error != nil {

		return nil, result.Error
	}
	return customer, nil

}

func (repo PostgresRepository) DeleteCustomer(id string) error {
	customer := domain.Customer{}
	if result := repo.DB.Find(&customer); result.Error != nil {
		return result.Error
	}
	fmt.Println(id)
	if err := repo.DB.Where("id = ? ", id).Delete(&customer).Error; err != nil {

		return err
	}
	return nil
}
