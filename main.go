package main

import (
	"github.com/yudhaananda/go-code-generator/handlers"
	"github.com/yudhaananda/go-code-generator/repositories"
	"github.com/yudhaananda/go-code-generator/services"
)

func main() {
	repo := repositories.Init(repositories.Param{})
	service := services.Init(services.Param{Repositories: repo})
	handler := handlers.Init(service)

	handler.Run()
}
