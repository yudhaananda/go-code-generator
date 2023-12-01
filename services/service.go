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
			ModelRepo:        param.Repositories.GenerateModels,
			MiddlewareRepo:   param.Repositories.GenerateMiddleware,
			FormatterRepo:    param.Repositories.GenerateFormatter,
			FilterRepo:       param.Repositories.GenerateFilter,
			RepositoriesRepo: param.Repositories.GenerateRepositories,
			ServicesRepo:     param.Repositories.GenerateServices,
			HandlerRepo:      param.Repositories.GenerateHandler,
			CmdRepo:          param.Repositories.GenerateCmd,
			MockRepo:         param.Repositories.GenerateMock,
			SqlRepo:          param.Repositories.GenerateSql,
			Zipping:          param.Repositories.Zipping,
		}),
	}
}
