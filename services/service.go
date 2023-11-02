package services

import (
	"github.com/yudhaananda/go-code-generator/repositories"
	generatemodels "github.com/yudhaananda/go-code-generator/services/generate_models"
)

type Services struct {
	GenerateModels generatemodels.Interface
}

type Param struct {
	Repositories *repositories.Repositories
}

func Init(param Param) *Services {
	return &Services{
		GenerateModels: generatemodels.Init(generatemodels.Params{Repo: param.Repositories.GenerateModels}),
	}
}
