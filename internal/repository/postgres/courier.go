package postgres

import (
	"flag"
	"fmt"

	"github.com/AntonyIS/go-foods/internal/core/domain"
)

var (
	tableName = flag.String("tableName", "courier", "Stores courier data ")
)

func (repo PostgresRepository) CreateCourier(courier *domain.Courier) (*domain.Courier, error) {

	if err := repo.DB.Create(&courier).Error; err != nil {
		return nil, err
	}
	return courier, nil
}
func (repo PostgresRepository) GetCouriers() (*[]domain.Courier, error) {

	courier := []domain.Courier{}
	if result := repo.DB.Find(&courier); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &courier, nil
}

func (repo PostgresRepository) GetCourier(id string) (*domain.Courier, error) {
	courier := domain.Courier{}
	if result := repo.DB.Find(&courier); result.Error != nil {
		return nil, result.Error
	}
	return &courier, nil
}

func (repo PostgresRepository) UpdateCourier(courier *domain.Courier) (*domain.Courier, error) {
	if result := repo.DB.Save(courier); result.Error != nil {

		return nil, result.Error
	}
	return courier, nil

}

func (repo PostgresRepository) DeleteCourier(id string) error {
	courier := domain.Courier{}
	if result := repo.DB.Find(&courier); result.Error != nil {
		return result.Error
	}
	fmt.Println(id)
	if err := repo.DB.Where("id = ? ", id).Delete(&courier).Error; err != nil {

		return err
	}
	return nil
}
