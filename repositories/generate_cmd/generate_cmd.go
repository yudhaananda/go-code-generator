package generatecmd

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	GenerateMain(project, folderName, fileName string, input models.GeneralTemplateInput) error
	GenerateMakefile(project, folderName, fileName string, input models.GeneralTemplateInput) error
	GenerateEnv(project, folderName, fileName string, input models.Env) error
}

type generateCmd struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateCmd{}
}

func (r *generateCmd) GenerateMain(project, folderName, fileName string, input models.GeneralTemplateInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/cmd/main.tmpl"
	t := template.Must(template.New("main.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateCmd) GenerateMakefile(project, folderName, fileName string, input models.GeneralTemplateInput) error {
	file, err := helper.CreateAdditionalFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/cmd/makefile.tmpl"
	t := template.Must(template.New("makefile.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateCmd) GenerateEnv(project, folderName, fileName string, input models.Env) error {
	file, err := helper.CreateAdditionalFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/cmd/env.tmpl"
	t := template.Must(template.New("env.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}
