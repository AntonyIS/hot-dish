package gin

import (
	"github.com/gin-gonic/gin"
)

type CourierHandler interface {
	PostsItem(*gin.Context)
	GetItems(*gin.Context)
	GetItem(*gin.Context)
	UpdateItem(*gin.Context)
	DeleteItem(*gin.Context)
}

type handler struct {
}

func NewCourierHandler() CourierHandler {
	return &handler{}
}
