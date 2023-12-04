package generate

import (
	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
	generateRepo "github.com/yudhaananda/go-code-generator/repositories/generate"
	"github.com/yudhaananda/go-code-generator/repositories/zipping"
)

type Interface interface {
	Generate(input models.Model) ([]byte, error)
}

type generate struct {
	generateRepo generateRepo.Interface
	zipping      zipping.Interface
}

type Params struct {
	GenerateRepo generateRepo.Interface
	Zipping      zipping.Interface
}

func Init(param Params) Interface {
	return &generate{
		generateRepo: param.GenerateRepo,
		zipping:      param.Zipping,
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
				TableName: helper.ConvertToSnakeCase(helper.ConvertToPlural(entity)),
			}
		)

		initRepo.Entity = append(initRepo.Entity, models.RepositoriesInitEntity{
			EntityName:                entity,
			EntityNameLowerCase:       helper.ConvertToLowerCase(entity),
			EntityNameSnakeCasePlural: helper.ConvertToSnakeCase(helper.ConvertToPlural(entity)),
			EntityNameSnakeCase:       helper.ConvertToSnakeCase(entity),
			ProjectName:               input.ProjectName,
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
		if err := s.generateRepo.Generate(
			input.ProjectName,
			"src/models",
			helper.ConvertToSnakeCase(entity)+".go",
			"templates/models/models.tmpl",
			"models.tmpl",
			models.CreateModelsInput{
				ProjectName: input.ProjectName,
				EntityName:  entity,
				EntityValue: entityValue,
			},
		); err != nil {
			return nil, err
		}

		// Generate Filter
		if err := s.generateRepo.Generate(
			input.ProjectName,
			"src/filter",
			helper.ConvertToSnakeCase(entity)+".go",
			"templates/filter/filter.tmpl",
			"filter.tmpl",
			models.FilterInput{
				EntityName:  entity,
				EntityValue: entityValue,
			},
		); err != nil {
			return nil, err
		}

		// Generate Repositories
		if err := s.generateRepo.Generate(
			input.ProjectName,
			"src/repositories/"+helper.ConvertToSnakeCase(entity),
			helper.ConvertToSnakeCase(entity)+".go",
			"templates/repositories/repositories.tmpl",
			"repositories.tmpl",
			models.RepositoriesInput{
				ProjectName:         input.ProjectName,
				EntityName:          entity,
				EntityNameCamelCase: helper.ConvertToCamelCase(entity),
				EntityNameLowerCase: helper.ConvertToLowerCase(entity),
			},
		); err != nil {
			return nil, err
		}

		// Generate Repositories Unit Test
		if err := s.generateRepo.Generate(
			input.ProjectName,
			"src/repositories/"+helper.ConvertToSnakeCase(entity),
			helper.ConvertToSnakeCase(entity)+"_test.go",
			"templates/repositories/repositories_test.tmpl",
			"repositories_test.tmpl",
			models.RepositoriesTestInput{
				ProjectName:         input.ProjectName,
				EntityName:          entity,
				EntityNameCamelCase: helper.ConvertToCamelCase(entity),
				EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
				MockTableMember:     mockTableMember,
				MockRow:             mockRow,
				WantMock:            wantMock,
				EntityNameLowerCase: helper.ConvertToLowerCase(entity),
			},
		); err != nil {
			return nil, err
		}

		// Generate Services
		if err := s.generateRepo.Generate(
			input.ProjectName,
			"src/services/"+helper.ConvertToSnakeCase(entity),
			helper.ConvertToSnakeCase(entity)+".go",
			"templates/services/services.tmpl",
			"services.tmpl",
			models.ServicesInput{
				ProjectName:         input.ProjectName,
				EntityNameLowerCase: helper.ConvertToLowerCase(entity),
				EntityName:          entity,
				EntityNameCamelCase: helper.ConvertToCamelCase(entity),
				EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
			},
		); err != nil {
			return nil, err
		}

		// Generate Services Unit Test
		if err := s.generateRepo.Generate(
			input.ProjectName,
			"src/services/"+helper.ConvertToSnakeCase(entity),
			helper.ConvertToSnakeCase(entity)+"_test.go",
			"templates/services/services_test.tmpl",
			"services_test.tmpl",
			models.ServicesTestInput{
				ProjectName:         input.ProjectName,
				EntityNameLowerCase: helper.ConvertToLowerCase(entity),
				EntityName:          entity,
				EntityNameCamelCase: helper.ConvertToCamelCase(entity),
				EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
			},
		); err != nil {
			return nil, err
		}

		// Generate Handler
		if err := s.generateRepo.Generate(
			input.ProjectName,
			"src/handler",
			helper.ConvertToSnakeCase(entity)+".go",
			"templates/handlers/handler.tmpl",
			"handler.tmpl",
			models.HandlerInput{
				ProjectName:         input.ProjectName,
				EntityNameLowerCase: helper.ConvertToLowerCase(entity),
				EntityName:          entity,
				EntityNameDash:      helper.ConvertToDash(entity),
			},
		); err != nil {
			return nil, err
		}

		// Generate Mock
		if err := s.generateRepo.Generate(
			input.ProjectName,
			"src/repositories/mock/"+helper.ConvertToSnakeCase(entity),
			helper.ConvertToSnakeCase(entity)+".go",
			"templates/mock/mock.tmpl",
			"mock.tmpl",
			models.MockInput{
				ProjectName:         input.ProjectName,
				EntityNameSnakeCase: helper.ConvertToSnakeCase(entity),
				EntityName:          entity,
			},
		); err != nil {
			return nil, err
		}
	}

	// Generate Model Helper
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/models/",
		"helper.go",
		"templates/models/helper.tmpl",
		"helper.tmpl",
		nil,
	); err != nil {
		return nil, err
	}

	// Generate Model ENV
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/models/",
		"env.go",
		"templates/models/env.tmpl",
		"env.tmpl",
		nil,
	); err != nil {
		return nil, err
	}

	// Generate Auth Model
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/models/",
		"auth.go",
		"templates/models/auth.tmpl",
		"auth.tmpl",
		nil,
	); err != nil {
		return nil, err
	}

	// Generate Auth Middleware
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/middleware/",
		"auth.go",
		"templates/middleware/auth.tmpl",
		"auth.tmpl",
		models.GeneralTemplateInput{
			ProjectName: input.ProjectName,
		},
	); err != nil {
		return nil, err
	}

	// Generate Auth Formatter
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/formatter/",
		"auth.go",
		"templates/formatter/auth.tmpl",
		"auth.tmpl",
		nil,
	); err != nil {
		return nil, err
	}

	// Generate Nullable Data Type Formatter
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/formatter/",
		"nullable_data_type.go",
		"templates/formatter/nullable_data_type.tmpl",
		"nullable_data_type.tmpl",
		nil,
	); err != nil {
		return nil, err
	}

	// Generate Paginated Items Formatter
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/formatter/",
		"paginated_items.go",
		"templates/formatter/paginated_items.tmpl",
		"paginated_items.tmpl",
		nil,
	); err != nil {
		return nil, err
	}

	// Generate Filter Helper
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/filter/",
		"helper.go",
		"templates/filter/helper.tmpl",
		"helper.tmpl",
		nil,
	); err != nil {
		return nil, err
	}

	// Generate Auth Repository
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/repositories/auth",
		"auth.go",
		"templates/repositories/auth.tmpl",
		"auth.tmpl",
		models.GeneralTemplateInput{
			ProjectName: input.ProjectName,
		},
	); err != nil {
		return nil, err
	}

	// Generate Base Repository
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/repositories/base",
		"base.go",
		"templates/repositories/base.tmpl",
		"base.tmpl",
		models.GeneralTemplateInput{
			ProjectName: input.ProjectName,
		},
	); err != nil {
		return nil, err
	}

	// Generate Base Query
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/repositories/base",
		"base_query.go",
		"templates/repositories/base_query.tmpl",
		"base_query.tmpl",
		nil,
	); err != nil {
		return nil, err
	}

	// Generate Initialize Repository
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/repositories",
		"repository.go",
		"templates/repositories/repositories_init.tmpl",
		"repositories_init.tmpl",
		initRepo,
	); err != nil {
		return nil, err
	}

	// Generate Auth Service
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/services/auth",
		"auth_test.go",
		"templates/services/auth_test.tmpl",
		"auth_test.tmpl",
		models.GeneralTemplateInput{
			ProjectName: input.ProjectName,
		},
	); err != nil {
		return nil, err
	}

	// Generate Auth Service Unit Test
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/services/auth",
		"auth.go",
		"templates/services/auth.tmpl",
		"auth.tmpl",
		models.GeneralTemplateInput{
			ProjectName: input.ProjectName,
		},
	); err != nil {
		return nil, err
	}

	// Generate Initialize Service
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/services",
		"service.go",
		"templates/services/services_init.tmpl",
		"services_init.tmpl",
		initService,
	); err != nil {
		return nil, err
	}

	// Generate Auth Handler
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/handler",
		"auth.go",
		"templates/handlers/auth.tmpl",
		"auth.tmpl",
		models.GeneralTemplateInput{
			ProjectName: input.ProjectName,
		},
	); err != nil {
		return nil, err
	}

	// Generate Mock Auth
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/repositories/mock/auth",
		"auth.go",
		"templates/mock/auth.tmpl",
		"auth.tmpl",
		nil,
	); err != nil {
		return nil, err
	}

	// Generate Initialize Handler
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/handler",
		"handler.go",
		"templates/handlers/handler_init.tmpl",
		"handler_init.tmpl",
		initHandler,
	); err != nil {
		return nil, err
	}

	// Generate Main
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"src/cmd",
		"main.go",
		"templates/cmd/main.tmpl",
		"main.tmpl",
		models.GeneralTemplateInput{
			ProjectName: input.ProjectName,
		},
	); err != nil {
		return nil, err
	}

	// Generate Makefile
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"",
		"Makefile",
		"templates/cmd/makefile.tmpl",
		"makefile.tmpl",
		models.GeneralTemplateInput{
			ProjectName: input.ProjectName,
		},
	); err != nil {
		return nil, err
	}

	// Generate Env
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"",
		".env",
		"templates/cmd/env.tmpl",
		"env.tmpl",
		models.Env{
			DBUser:     input.Database.User,
			DBPassword: input.Database.Pass,
			DBPort:     input.Database.Port,
			DBHost:     input.Database.Host,
			DBName:     input.Database.Name,
			JwtToken:   input.Database.JwtToken,
			DBType:     input.Database.Type,
		},
	); err != nil {
		return nil, err
	}

	// Generate Initialize Sql
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"docs/sql",
		"0000000000_init_db.sql",
		"templates/sql/sql.tmpl",
		"sql.tmpl",
		sql,
	); err != nil {
		return nil, err
	}

	// Generate Config.yaml
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"docs/sql",
		"config.yaml",
		"templates/sql/config.tmpl",
		"config.tmpl",
		models.Env{
			DBUser:     input.Database.User,
			DBPassword: input.Database.Pass,
			DBPort:     input.Database.Port,
			DBHost:     input.Database.Host,
			DBName:     input.Database.Name,
			JwtToken:   input.Database.JwtToken,
			DBType:     input.Database.Type,
		},
	); err != nil {
		return nil, err
	}

	// Generate Init.sh
	if err := s.generateRepo.Generate(
		input.ProjectName,
		"docs/sql",
		"init.sh",
		"templates/sql/init.tmpl",
		"init.tmpl",
		models.Env{
			DBUser:     input.Database.User,
			DBPassword: input.Database.Pass,
			DBPort:     input.Database.Port,
			DBHost:     input.Database.Host,
			DBName:     input.Database.Name,
			JwtToken:   input.Database.JwtToken,
			DBType:     input.Database.Type,
		},
	); err != nil {
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
