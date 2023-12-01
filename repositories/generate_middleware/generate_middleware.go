package generatemiddleware

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	GenerateAuthMiddleware(string, string, string, models.GeneralTemplateInput) error
}

type generateMiddleware struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateMiddleware{}
}

func (r *generateMiddleware) GenerateAuthMiddleware(project, folderName, fileName string, input models.GeneralTemplateInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/middleware/auth.tmpl"
	t := template.Must(template.New("auth.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}
