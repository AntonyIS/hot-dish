package ports

import "github.com/AntonyIS/go-foods/internal/core/domain"

type RestaurantRepository interface {
	CreateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error)
	GetRestaurants() (*[]domain.Restaurant, error)
	GetRestaurant(id string) (*domain.Restaurant, error)
	UpdateRestaurant(restaurant *domain.Restaurant) (*domain.Restaurant, error)
	DeleteRestaurant(id string) error
}

type CustomerRepository interface {
	CreateCustomer(customer *domain.Customer) (*domain.Customer, error)
	GetCustomers() (*[]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, error)
	UpdateCustomer(customer *domain.Customer) (*domain.Customer, error)
	DeleteCustomer(id string) error
}

type CourierRepository interface {
	CreateCourier(courier *domain.Courier) (*domain.Courier, error)
	GetCouriers() (*[]domain.Courier, error)
	GetCourier(id string) (*domain.Courier, error)
	UpdateCourier(courier *domain.Courier) (*domain.Courier, error)
	DeleteCourier(id string) error
}

type ItemRepository interface {
	CourierRepository
	CustomerRepository
	RestaurantRepository
}
