package generatemodels

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	CreateModels(string, string, string, models.CreateModelsInput) error
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

	var tmplFile = "templates/models.tmpl"
	t := template.Must(template.New("models.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		panic(err)
	}
	return nil
}
