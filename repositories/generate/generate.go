package generate

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
)

type Interface interface {
	Generate(project, folderName, fileName, tmplPath, tmplFileName string, input any) error
}

type generate struct {
}

func Init() Interface {
	return &generate{}
}

func (r *generate) Generate(project, folderName, fileName, tmplPath, tmplFileName string, input any) error {
	file, err := helper.CreateAdditionalFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = tmplPath
	t := template.Must(template.New(tmplFileName).ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}
