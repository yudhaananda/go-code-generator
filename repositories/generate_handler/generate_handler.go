package generatehandler

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	GenerateAuthHandler(project, folderName, fileName string, input models.GeneralTemplateInput) error
	GenerateHandler(project, folderName, fileName string, input models.HandlerInput) error
	GenerateHandlerInit(project, folderName, fileName string, input models.HandlerInitInput) error
}

type generateHandler struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateHandler{}
}

func (r *generateHandler) GenerateAuthHandler(project, folderName, fileName string, input models.GeneralTemplateInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/handlers/auth.tmpl"
	t := template.Must(template.New("auth.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateHandler) GenerateHandler(project, folderName, fileName string, input models.HandlerInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/handlers/handler.tmpl"
	t := template.Must(template.New("handler.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateHandler) GenerateHandlerInit(project, folderName, fileName string, input models.HandlerInitInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/handlers/handler_init.tmpl"
	t := template.Must(template.New("handler_init.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}
