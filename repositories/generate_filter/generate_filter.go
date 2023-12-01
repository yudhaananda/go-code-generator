package generatefilter

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	GenerateFilterHelper(string, string, string) error
	GenerateFilter(string, string, string, models.FilterInput) error
}

type generateFilter struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateFilter{}
}

func (*generateFilter) GenerateFilter(project, folderName, fileName string, input models.FilterInput) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/filter/filter.tmpl"
	t := template.Must(template.New("filter.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateFilter) GenerateFilterHelper(project, folderName, fileName string) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/filter/helper.tmpl"
	t := template.Must(template.New("helper.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}
