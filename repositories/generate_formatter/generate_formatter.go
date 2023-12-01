package generateformatter

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
)

type Interface interface {
	GenerateAuthFormatter(string, string, string) error
	GenerateNullableDataTypeFormatter(string, string, string) error
	GeneratePaginatedItemsFormatter(string, string, string) error
}

type generateFormatter struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateFormatter{}
}

func (r *generateFormatter) GenerateAuthFormatter(project, folderName, fileName string) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/formatter/auth.tmpl"
	t := template.Must(template.New("auth.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateFormatter) GenerateNullableDataTypeFormatter(project, folderName, fileName string) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/formatter/nullable_data_type.tmpl"
	t := template.Must(template.New("nullable_data_type.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateFormatter) GeneratePaginatedItemsFormatter(project, folderName, fileName string) error {
	file, err := helper.CreateFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/formatter/paginated_items.tmpl"
	t := template.Must(template.New("paginated_items.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}
