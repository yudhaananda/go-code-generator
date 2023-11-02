package generatemodels

import (
	"github.com/yudhaananda/go-code-generator/models"
	generatemodels "github.com/yudhaananda/go-code-generator/repositories/generate_models"
)

type Interface interface {
	CreateModels(input models.Model) error
}

type generateModels struct {
	Repo generatemodels.Interface
}

type Params struct {
	Repo generatemodels.Interface
}

func Init(param Params) Interface {
	return &generateModels{
		Repo: param.Repo,
	}
}

func (s *generateModels) CreateModels(input models.Model) error {
	return nil
}
