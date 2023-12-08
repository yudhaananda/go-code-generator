package main

import (
	"generator/handlers"
	"generator/repositories"
	"generator/services"
)

func main() {
	repo := repositories.Init(repositories.Param{})
	service := services.Init(services.Param{Repositories: repo})
	handler := handlers.Init(service)

	if err := handler.Generate(); err != nil {
		panic(err)
	}
}
