package main

import (
	"go-restapi-template/app/infrastructure/db"
	"go-restapi-template/app/infrastructure/router"
	"go-restapi-template/app/interfaces/repository"
	"go-restapi-template/app/interfaces/controller"
	"go-restapi-template/app/domain/usecase"

)

func main() {
	db := db.NewDB()
	if db == nil {
		return
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	route := router.NewRouter(userController)
	route.Run(":8080")
}