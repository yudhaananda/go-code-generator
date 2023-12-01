package generateservices

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	GenerateAuthServices(string, string, string, models.GeneralTemplateInput) error
	GenerateAuthTestServices(string, string, string, models.GeneralTemplateInput) error
	GenerateServices(string, string, string, models.ServicesInput) error
	GenerateTestServices(string, string, string, models.ServicesTestInput) error
	GenerateInitServices(string, string, string, models.ServicesInitInput) error
}

type generateServices struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateServices{}
}

func (r *generateServices) GenerateAuthServices(project, folderName, fileName string, input models.GeneralTemplateInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/services/auth.tmpl"
	t := template.Must(template.New("auth.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateServices) GenerateAuthTestServices(project, folderName, fileName string, input models.GeneralTemplateInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/services/auth_test.tmpl"
	t := template.Must(template.New("auth_test.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateServices) GenerateServices(project, folderName, fileName string, input models.ServicesInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/services/services.tmpl"
	t := template.Must(template.New("services.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateServices) GenerateTestServices(project, folderName, fileName string, input models.ServicesTestInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/services/services_test.tmpl"
	t := template.Must(template.New("services_test.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateServices) GenerateInitServices(project, folderName, fileName string, input models.ServicesInitInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/services/services_init.tmpl"
	t := template.Must(template.New("services_init.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}
