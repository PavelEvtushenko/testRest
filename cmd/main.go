package main

import (
	"log"
	todo "restApi"
	"restApi/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("ошибка сервера 1 %s", err.Error())
	}

}
