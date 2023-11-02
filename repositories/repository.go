package repositories

import (
	"github.com/yudhaananda/go-code-generator/models"
	generatemodels "github.com/yudhaananda/go-code-generator/repositories/generate_models"
)

type Repositories struct {
	GenerateModels generatemodels.Interface
}

type Param struct {
	EntityName  string
	EntityValue []models.EntityValue
}

func Init(param Param) *Repositories {
	return &Repositories{
		GenerateModels: generatemodels.Init(generatemodels.Params{}),
	}
}
