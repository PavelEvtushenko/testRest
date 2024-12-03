package main

import (
	"fmt"
	"log"
	todo "restApi"
	"restApi/pkg/handler"
	"restApi/pkg/repository"
	"restApi/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("ошибка сервера 1 %s", err.Error())
	}

	fmt.Println("test")
}
