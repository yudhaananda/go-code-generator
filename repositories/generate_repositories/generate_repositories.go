package generaterepositories

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	GenerateAuthRepositories(string, string, string, models.GeneralTemplateInput) error
	GenerateBaseRepositories(string, string, string, models.GeneralTemplateInput) error
	GenerateBaseQueryRepositories(string, string, string) error
	GenerateRepositories(string, string, string, models.RepositoriesInput) error
	GenerateTestRepositories(string, string, string, models.RepositoriesTestInput) error
	GenerateInitRepositories(string, string, string, models.RepositoriesInitInput) error
}

type generateRepositories struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateRepositories{}
}

func (r *generateRepositories) GenerateAuthRepositories(project, folderName, fileName string, input models.GeneralTemplateInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/repositories/auth.tmpl"
	t := template.Must(template.New("auth.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateRepositories) GenerateBaseRepositories(project, folderName, fileName string, input models.GeneralTemplateInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/repositories/base.tmpl"
	t := template.Must(template.New("base.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateRepositories) GenerateBaseQueryRepositories(project, folderName, fileName string) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/repositories/base_query.tmpl"
	t := template.Must(template.New("base_query.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateRepositories) GenerateRepositories(project, folderName, fileName string, input models.RepositoriesInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/repositories/repositories.tmpl"
	t := template.Must(template.New("repositories.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateRepositories) GenerateTestRepositories(project, folderName, fileName string, input models.RepositoriesTestInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/repositories/repositories_test.tmpl"
	t := template.Must(template.New("repositories_test.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateRepositories) GenerateInitRepositories(project, folderName, fileName string, input models.RepositoriesInitInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/repositories/repositories_init.tmpl"
	t := template.Must(template.New("repositories_init.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}
