package restaurant

import (
	"github.com/AntonyIS/go-foods/internal/core/domain"
	"github.com/AntonyIS/go-foods/internal/core/ports"
)

type restaurantService struct {
	repo ports.RestaurantRepository
}

func NewRestaurantService(repo ports.RestaurantRepository) ports.RestaurantService {
	return &restaurantService{
		repo,
	}
}

func (cs restaurantService) CreateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error) {
	return cs.repo.CreateRestaurant(restaurant)
}

func (cs restaurantService) GetRestaurants() (*[]domain.Restaurant, error) {
	return cs.repo.GetRestaurants()
}

func (cs restaurantService) GetRestaurant(id string) (*domain.Restaurant, error) {
	return cs.repo.GetRestaurant(id)
}

func (cs restaurantService) UpdateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error) {
	return cs.repo.UpdateRestaurant(restaurant)
}

func (cs restaurantService) DeleteRestaurant(id string) error {
	return cs.repo.DeleteRestaurant(id)
}
