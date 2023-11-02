package main

import (
	"github.com/yudhaananda/go-code-generator/models"
	"github.com/yudhaananda/go-code-generator/repositories"
)

func main() {
	repo := repositories.Init(repositories.Param{})
	repo.GenerateModels.CreateModels("test", "test", "test.go", models.CreateModelsInput{
		EntityName: "test",
		EntityValue: []models.EntityValue{
			{
				EntityValueName:          "test",
				EntityDataType:           "test",
				EntityValueNameSnakeCase: "test",
				EntityValueNameCamelCase: "test",
			},
		},
	})
}
