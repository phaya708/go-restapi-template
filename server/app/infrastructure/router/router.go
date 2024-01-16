package router

import (
	"github.com/gin-gonic/gin"
	"go-restapi-template/app/interfaces/controller"
)

func NewRouter(uc controller.UserController) *gin.Engine {
	route := gin.Default()
	route.GET("/users", uc.GetAll)
	route.GET("/users/:id", uc.GetByID)
	route.POST("/users", uc.Create)
	route.PUT("/users/:id", uc.Update)
	route.DELETE("/users/:id", uc.Delete)
	return route
}
