package generate

import (
	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
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

type Interface interface {
	Generate(input models.Model) ([]byte, error)
}

type generate struct {
	modelRepo        generatemodels.Interface
	middlewareRepo   generatemiddleware.Interface
	formatterRepo    generateformatter.Interface
	filterRepo       generatefilter.Interface
	repositoriesRepo generaterepositories.Interface
	servicesRepo     generateservices.Interface
	handlerRepo      generatehandler.Interface
	cmdRepo          generatecmd.Interface
	mockRepo         generatemock.Interface
	sqlRepo          generatesql.Interface
	zipping          zipping.Interface
}

type Params struct {
	ModelRepo        generatemodels.Interface
	MiddlewareRepo   generatemiddleware.Interface
	FormatterRepo    generateformatter.Interface
	FilterRepo       generatefilter.Interface
	RepositoriesRepo generaterepositories.Interface
	ServicesRepo     generateservices.Interface
	HandlerRepo      generatehandler.Interface
	CmdRepo          generatecmd.Interface
	MockRepo         generatemock.Interface
	SqlRepo          generatesql.Interface
	Zipping          zipping.Interface
}

func Init(param Params) Interface {
	return &generate{
		modelRepo:        param.ModelRepo,
		middlewareRepo:   param.MiddlewareRepo,
		formatterRepo:    param.FormatterRepo,
		filterRepo:       param.FilterRepo,
		repositoriesRepo: param.RepositoriesRepo,
		servicesRepo:     param.ServicesRepo,
		handlerRepo:      param.HandlerRepo,
		cmdRepo:          param.CmdRepo,
		mockRepo:         param.MockRepo,
		sqlRepo:          param.SqlRepo,
		zipping:          param.Zipping,
	}
}

func (s *generate) Generate(input models.Model) ([]byte, error) {
	var (
		initRepo = models.RepositoriesInitInput{
			ProjectName: input.ProjectName,
		}
		initService = models.ServicesInitInput{
			ProjectName: input.ProjectName,
		}
		initHandler = models.HandlerInitInput{
			ProjectName: input.ProjectName,
		}
		sql = models.Table{}
	)
	for entity, members := range input.Entity {
		var (
			entityValue     = []models.EntityValue{}
			mockTableMember string
			mockRow         string
			wantMock        string
			tableMember     = models.TableMember{
				TableName: helper.ConvertToSnakeCase(entity),
			}
		)

		initRepo.Entity = append(initRepo.Entity, models.RepositoriesInitEntity{
			EntityName:          entity,
			EntityNameLowerCase: helper.ConvertToLowerCase(entity),
			EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
			ProjectName:         input.ProjectName,
		})

		initService.Entity = append(initService.Entity, models.ServicesInitEntity{
			EntityName:          entity,
			EntityNameLowerCase: helper.ConvertToLowerCase(entity),
			EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
			ProjectName:         initRepo.ProjectName,
		})

		initHandler.Entity = append(initHandler.Entity, models.HanlderInitEntity{
			EntityNameLowerCase: helper.ConvertToLowerCase(entity),
			EntityName:          entity,
			EntityNameDash:      helper.ConvertToDash(entity),
		})

		for _, member := range members {
			item := member.GetTableValue()
			entityValue = append(entityValue, member.GetEntityValue())
			tableMember.TableItems = append(tableMember.TableItems, models.TableItem{Item: item})
			mockTableMember += member.GetMockTableMember()
			mockRow += member.GetMockRow()
			wantMock += member.GetWantMock()
		}

		sql.Entity = append(sql.Entity, tableMember)

		// Generate Models
		if err := s.modelRepo.CreateModels(input.ProjectName, "models", helper.ConvertToSnakeCase(entity)+".go", models.CreateModelsInput{
			ProjectName: input.ProjectName,
			EntityName:  entity,
			EntityValue: entityValue,
		}); err != nil {
			return nil, err
		}

		// Generate Filter
		if err := s.filterRepo.GenerateFilter(input.ProjectName, "filter", helper.ConvertToSnakeCase(entity)+".go", models.FilterInput{
			EntityName:  entity,
			EntityValue: entityValue,
		}); err != nil {
			return nil, err
		}

		// Generate Repositories
		if err := s.repositoriesRepo.GenerateRepositories(input.ProjectName, "repositories/"+helper.ConvertToSnakeCase(entity), helper.ConvertToSnakeCase(entity)+".go", models.RepositoriesInput{
			ProjectName:         input.ProjectName,
			EntityName:          entity,
			EntityNameCamelCase: helper.ConvertToCamelCase(entity),
			EntityNameLowerCase: helper.ConvertToLowerCase(entity),
		}); err != nil {
			return nil, err
		}

		// Generate Repositories Unit Test
		if err := s.repositoriesRepo.GenerateTestRepositories(input.ProjectName, "repositories/"+helper.ConvertToSnakeCase(entity), helper.ConvertToSnakeCase(entity)+"_test.go", models.RepositoriesTestInput{
			ProjectName:         input.ProjectName,
			EntityName:          entity,
			EntityNameCamelCase: helper.ConvertToCamelCase(entity),
			EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
			MockTableMember:     mockTableMember,
			MockRow:             mockRow,
			WantMock:            wantMock,
			EntityNameLowerCase: helper.ConvertToLowerCase(entity),
		}); err != nil {
			return nil, err
		}

		// Generate Services
		if err := s.servicesRepo.GenerateServices(input.ProjectName, "services/"+helper.ConvertToSnakeCase(entity), helper.ConvertToSnakeCase(entity)+".go", models.ServicesInput{
			ProjectName:         input.ProjectName,
			EntityNameLowerCase: helper.ConvertToLowerCase(entity),
			EntityName:          entity,
			EntityNameCamelCase: helper.ConvertToCamelCase(entity),
			EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
		}); err != nil {
			return nil, err
		}

		// Generate Services Unit Test
		if err := s.servicesRepo.GenerateTestServices(input.ProjectName, "services/"+helper.ConvertToSnakeCase(entity), helper.ConvertToSnakeCase(entity)+"_test.go", models.ServicesTestInput{
			ProjectName:         input.ProjectName,
			EntityNameLowerCase: helper.ConvertToLowerCase(entity),
			EntityName:          entity,
			EntityNameCamelCase: helper.ConvertToCamelCase(entity),
			EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
		}); err != nil {
			return nil, err
		}

		// Generate Handler
		if err := s.handlerRepo.GenerateHandler(input.ProjectName, "handler/", helper.ConvertToSnakeCase(entity)+".go", models.HandlerInput{
			ProjectName:         input.ProjectName,
			EntityNameLowerCase: helper.ConvertToLowerCase(entity),
			EntityName:          entity,
			EntityNameDash:      helper.ConvertToDash(entity),
		}); err != nil {
			return nil, err
		}

		// Generate Mock
		if err := s.mockRepo.GenerateMock(input.ProjectName, "repositories/mock/"+helper.ConvertToSnakeCase(entity), helper.ConvertToSnakeCase(entity)+".go", models.MockInput{
			ProjectName:         input.ProjectName,
			EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
			EntityName:          entity,
		}); err != nil {
			return nil, err
		}
	}

	// Generate Model Helper
	if err := s.modelRepo.GenerateModelHelper(input.ProjectName, "models", "helper.go"); err != nil {
		return nil, err
	}

	// Generate Model ENV
	if err := s.modelRepo.GenerateModelEnv(input.ProjectName, "models", "env.go"); err != nil {
		return nil, err
	}

	// Generate Auth Model
	if err := s.modelRepo.GenerateModelAuth(input.ProjectName, "models", "auth.go"); err != nil {
		return nil, err
	}

	// Generate Auth Middleware
	if err := s.middlewareRepo.GenerateAuthMiddleware(input.ProjectName, "middleware", "auth.go", models.GeneralTemplateInput{
		ProjectName: input.ProjectName,
	}); err != nil {
		return nil, err
	}

	// Generate Auth Formatter
	if err := s.formatterRepo.GenerateAuthFormatter(input.ProjectName, "formatter", "auth.go"); err != nil {
		return nil, err
	}

	// Generate Nullable Data Type Formatter
	if err := s.formatterRepo.GenerateNullableDataTypeFormatter(input.ProjectName, "formatter", "nullable_data_type.go"); err != nil {
		return nil, err
	}

	// Generate Paginated Items Formatter
	if err := s.formatterRepo.GeneratePaginatedItemsFormatter(input.ProjectName, "formatter", "paginated_items.go"); err != nil {
		return nil, err
	}

	// Generate Filter Helper
	if err := s.filterRepo.GenerateFilterHelper(input.ProjectName, "filter", "helper.go"); err != nil {
		return nil, err
	}

	// Generate Auth Repository
	if err := s.repositoriesRepo.GenerateAuthRepositories(input.ProjectName, "repositories/auth", "auth.go", models.GeneralTemplateInput{
		ProjectName: input.ProjectName,
	}); err != nil {
		return nil, err
	}

	// Generate Base Repository
	if err := s.repositoriesRepo.GenerateBaseRepositories(input.ProjectName, "repositories/base", "base.go", models.GeneralTemplateInput{
		ProjectName: input.ProjectName,
	}); err != nil {
		return nil, err
	}

	// Generate Base Query
	if err := s.repositoriesRepo.GenerateBaseQueryRepositories(input.ProjectName, "repositories/base", "base_query.go"); err != nil {
		return nil, err
	}

	// Generate Initialize Repository
	if err := s.repositoriesRepo.GenerateInitRepositories(input.ProjectName, "repositories", "repository.go", initRepo); err != nil {
		return nil, err
	}

	// Generate Auth Service
	if err := s.servicesRepo.GenerateAuthServices(input.ProjectName, "services/auth", "auth.go", models.GeneralTemplateInput{
		ProjectName: input.ProjectName,
	}); err != nil {
		return nil, err
	}

	// Generate Auth Service Unit Test
	if err := s.servicesRepo.GenerateAuthTestServices(input.ProjectName, "services/auth", "auth_test.go", models.GeneralTemplateInput{
		ProjectName: input.ProjectName,
	}); err != nil {
		return nil, err
	}

	// Generate Initialize Service
	if err := s.servicesRepo.GenerateInitServices(input.ProjectName, "services", "service.go", initService); err != nil {
		return nil, err
	}

	// Generate Auth Handler
	if err := s.handlerRepo.GenerateAuthHandler(input.ProjectName, "handler", "auth.go", models.GeneralTemplateInput{
		ProjectName: input.ProjectName,
	}); err != nil {
		return nil, err
	}

	// Generate Mock Auth
	if err := s.mockRepo.GenerateAuthMock(input.ProjectName, "repositories/mock/auth", "auth.go"); err != nil {
		return nil, err
	}

	// Generate Initialize Handler
	if err := s.handlerRepo.GenerateHandlerInit(input.ProjectName, "handler", "handler.go", initHandler); err != nil {
		return nil, err
	}

	// Generate Main
	if err := s.cmdRepo.GenerateMain(input.ProjectName, "cmd", "main.go", models.GeneralTemplateInput{
		ProjectName: input.ProjectName,
	}); err != nil {
		return nil, err
	}

	// Generate Makefile
	if err := s.cmdRepo.GenerateMakefile(input.ProjectName, "", "Makefile", models.GeneralTemplateInput{
		ProjectName: input.ProjectName,
	}); err != nil {
		return nil, err
	}

	// Generate Env
	if err := s.cmdRepo.GenerateEnv(input.ProjectName, "", ".env", models.Env{
		DBUser:     input.Database.User,
		DBPassword: input.Database.Pass,
		DBPort:     input.Database.Port,
		DBHost:     input.Database.Host,
		DBName:     input.Database.Name,
		JwtToken:   input.Database.JwtToken,
		DBType:     input.Database.Type,
	}); err != nil {
		return nil, err
	}

	// Generate Initialize Sql
	if err := s.sqlRepo.CreateSql(input.ProjectName, "docs/sql", "0000000000_init_db.sql", sql); err != nil {
		return nil, err
	}

	// Generate Config.yaml
	if err := s.sqlRepo.CreateConfig(input.ProjectName, "docs/sql", "config.yaml", models.Env{
		DBUser:     input.Database.User,
		DBPassword: input.Database.Pass,
		DBPort:     input.Database.Port,
		DBHost:     input.Database.Host,
		DBName:     input.Database.Name,
		JwtToken:   input.Database.JwtToken,
		DBType:     input.Database.Type,
	}); err != nil {
		return nil, err
	}

	// Generate Init.sh
	if err := s.sqlRepo.CreateInit(input.ProjectName, "docs/sql", "init.sh", models.Env{
		DBUser:     input.Database.User,
		DBPassword: input.Database.Pass,
		DBPort:     input.Database.Port,
		DBHost:     input.Database.Host,
		DBName:     input.Database.Name,
		JwtToken:   input.Database.JwtToken,
		DBType:     input.Database.Type,
	}); err != nil {
		return nil, err
	}

	// Zipping
	result, err := s.zipping.Zipping(input.ProjectName)
	if err != nil {
		return nil, err
	}
	if err := s.zipping.Delete(input.ProjectName); err != nil {
		return nil, err
	}
	return result, nil
}
