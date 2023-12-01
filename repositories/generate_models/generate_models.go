package generatemodels

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	CreateModels(string, string, string, models.CreateModelsInput) error
	GenerateModelHelper(string, string, string) error
	GenerateModelEnv(string, string, string) error
	GenerateModelAuth(string, string, string) error
}

type generateModels struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateModels{}
}

func (r *generateModels) CreateModels(project, folderName, fileName string, input models.CreateModelsInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/models/models.tmpl"
	t := template.Must(template.New("models.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateModels) GenerateModelHelper(project, folderName, fileName string) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/models/helper.tmpl"
	t := template.Must(template.New("helper.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateModels) GenerateModelEnv(project, folderName, fileName string) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/models/env.tmpl"
	t := template.Must(template.New("env.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateModels) GenerateModelAuth(project, folderName, fileName string) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/models/auth.tmpl"
	t := template.Must(template.New("auth.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}
