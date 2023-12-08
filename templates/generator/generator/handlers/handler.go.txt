package handlers

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"generator/models"
	"generator/services"
)

type handler struct {
	Service *services.Services
}

func Init(service *services.Services) *handler {
	return &handler{
		Service: service,
	}
}

func (h *handler) Generate() error {

	file, err := os.Open("project.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	var jsonContent models.Model
	err = json.Unmarshal(data, &jsonContent)

	if err != nil {
		return err
	}

	err = h.Service.Generate.Generate(jsonContent)
	if err != nil {
		return err
	}
	return nil
}
