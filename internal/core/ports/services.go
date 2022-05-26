package ports

import "github.com/AntonyIS/go-foods/internal/core/domain"

type RestaurantService interface {
	CreateRestaurant(r *domain.Restaurant) (*domain.Restaurant, error)
	GetRestaurants() (*[]domain.Restaurant, error)
	GetRestaurant(id string) (*domain.Restaurant, error)
	UpdateRestaurant(r *domain.Restaurant) (*domain.Restaurant, error)
	DeleteRestaurant(id string) (*domain.Restaurant, error)
}

type CustomerService interface {
	CreateCustomer(c *domain.Customer) (*domain.Customer, error)
	GetCustomers() (*[]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, error)
	UpdateCustomer(c *domain.Customer) (*domain.Customer, error)
	DeleteCustomer(id string) (*domain.Customer, error)
}

type CourierService interface {
	CreateCourier(r *domain.Courier) (*domain.Courier, error)
	GetCouriers() (*[]domain.Courier, error)
	GetCourier(id string) (*domain.Courier, error)
	UpdateCourier(r *domain.Courier) (*domain.Courier, error)
	DeleteCourier(id string) (*domain.Courier, error)
}
