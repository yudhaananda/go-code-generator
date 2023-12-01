package repositories

import (
	generatecmd "github.com/yudhaananda/go-code-generator/repositories/generate_cmd"
	generatefilter "github.com/yudhaananda/go-code-generator/repositories/generate_filter"
	generateformatter "github.com/yudhaananda/go-code-generator/repositories/generate_formatter"
	generatehandler "github.com/yudhaananda/go-code-generator/repositories/generate_handler"
	generatemiddleware "github.com/yudhaananda/go-code-generator/repositories/generate_middleware"
	generatemock "github.com/yudhaananda/go-code-generator/repositories/generate_mock"
	generatemodels "github.com/yudhaananda/go-code-generator/repositories/generate_models"
	generaterepositories "github.com/yudhaananda/go-code-generator/repositories/generate_repositories"
	generateservices "github.com/yudhaananda/go-code-generator/repositories/generate_services"
	generatesql "github.com/yudhaananda/go-code-generator/repositories/generate_sql"
	"github.com/yudhaananda/go-code-generator/repositories/zipping"
)

type Repositories struct {
	GenerateModels       generatemodels.Interface
	GenerateMiddleware   generatemiddleware.Interface
	GenerateFormatter    generateformatter.Interface
	GenerateFilter       generatefilter.Interface
	GenerateRepositories generaterepositories.Interface
	GenerateServices     generateservices.Interface
	GenerateHandler      generatehandler.Interface
	GenerateCmd          generatecmd.Interface
	GenerateMock         generatemock.Interface
	GenerateSql          generatesql.Interface
	Zipping              zipping.Interface
}

type Param struct {
}

func Init(param Param) *Repositories {
	return &Repositories{
		GenerateModels:       generatemodels.Init(generatemodels.Params{}),
		GenerateMiddleware:   generatemiddleware.Init(generatemiddleware.Params{}),
		GenerateFormatter:    generateformatter.Init(generateformatter.Params{}),
		GenerateFilter:       generatefilter.Init(generatefilter.Params{}),
		GenerateRepositories: generaterepositories.Init(generaterepositories.Params{}),
		GenerateServices:     generateservices.Init(generateservices.Params{}),
		GenerateHandler:      generatehandler.Init(generatehandler.Params{}),
		GenerateCmd:          generatecmd.Init(generatecmd.Params{}),
		GenerateMock:         generatemock.Init(generatemock.Params{}),
		GenerateSql:          generatesql.Init(generatesql.Params{}),
		Zipping:              zipping.Init(),
	}
}
