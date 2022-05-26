package gin

import (
	"net/http"

	"github.com/AntonyIS/go-foods/internal/core/domain"
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

func (handler) CreateCourier(ctx gin.Context) {
	courier := domain.Courier{}

	if err := ctx.ShouldBindJSON(&courier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
