package generatemock

import (
	"html/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	GenerateAuthMock(project, folderName, fileName string) error
	GenerateMock(project, folderName, fileName string, input models.MockInput) error
}

type generateMock struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateMock{}
}

func (r *generateMock) GenerateAuthMock(project, folderName, fileName string) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/mock/auth.tmpl"
	t := template.Must(template.New("auth.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateMock) GenerateMock(project, folderName, fileName string, input models.MockInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/mock/mock.tmpl"
	t := template.Must(template.New("mock.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}
