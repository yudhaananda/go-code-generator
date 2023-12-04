package services

import (
	"github.com/yudhaananda/go-code-generator/repositories"
	"github.com/yudhaananda/go-code-generator/services/generate"
)

type Services struct {
	Generate generate.Interface
}

type Param struct {
	Repositories *repositories.Repositories
}

func Init(param Param) *Services {
	return &Services{
		Generate: generate.Init(generate.Params{
			GenerateRepo: param.Repositories.Generate,
			Zipping:      param.Repositories.Zipping,
		}),
	}
}
