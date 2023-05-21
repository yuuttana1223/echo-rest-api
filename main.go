package main

import (
	"echo-rest-api/controller"
	"echo-rest-api/db"
	"echo-rest-api/repository"
	"echo-rest-api/router"
	"echo-rest-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
