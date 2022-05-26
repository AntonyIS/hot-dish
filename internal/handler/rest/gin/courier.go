package gin

import (
	"github.com/gin-gonic/gin"
)

type CourierHandler interface {
	CreateRestaurant(*gin.Context)
	GetRestaurants(*gin.Context)
	GetRestaurant(*gin.Context)
	UpdateRestaurant(*gin.Context)
	DeleteRestaurant(*gin.Context)
}

type handler struct {
}

func NewCourierHandler() CourierHandler {
	return &handler{}
}
