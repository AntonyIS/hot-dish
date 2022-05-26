package postgres

import (
	"fmt"

	"github.com/AntonyIS/go-foods/internal/core/domain"
)

func (repo PostgresRepository) CreateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error) {

	if err := repo.DB.Create(&restaurant).Error; err != nil {
		return nil, err
	}
	return restaurant, nil
}
func (repo PostgresRepository) GetRestaurants() (*[]domain.Restaurant, error) {

	restaurant := []domain.Restaurant{}
	if result := repo.DB.Find(&restaurant); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &restaurant, nil
}

func (repo PostgresRepository) GetRestaurant(id string) (*domain.Restaurant, error) {
	restaurant := domain.Restaurant{}
	if result := repo.DB.Find(&restaurant); result.Error != nil {
		return nil, result.Error
	}
	return &restaurant, nil
}

func (repo PostgresRepository) UpdateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error) {
	if result := repo.DB.Save(restaurant); result.Error != nil {

		return nil, result.Error
	}
	return restaurant, nil

}

func (repo PostgresRepository) DeleteRestaurant(id string) error {
	restaurant := domain.Restaurant{}
	if result := repo.DB.Find(&restaurant); result.Error != nil {
		return result.Error
	}

	if err := repo.DB.Where("id = ? ", id).Delete(&restaurant).Error; err != nil {

		return err
	}
	return nil
}
