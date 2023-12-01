package generatesql

import (
	"text/template"

	"github.com/yudhaananda/go-code-generator/helper"
	"github.com/yudhaananda/go-code-generator/models"
)

type Interface interface {
	CreateSql(project, folderName, fileName string, input models.Table) error
	CreateInit(project, folderName, fileName string, input models.Env) error
	CreateConfig(project, folderName, fileName string, input models.Env) error
}

type generateSql struct {
}

type Params struct {
}

func Init(param Params) Interface {
	return &generateSql{}
}

func (r *generateSql) CreateSql(project, folderName, fileName string, input models.Table) error {
	file, err := helper.CreateAdditionalFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/sql/sql.tmpl"
	t := template.Must(template.New("sql.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateSql) CreateConfig(project, folderName, fileName string, input models.Env) error {
	file, err := helper.CreateAdditionalFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/sql/config.tmpl"
	t := template.Must(template.New("config.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *generateSql) CreateInit(project, folderName, fileName string, input models.Env) error {
	file, err := helper.CreateAdditionalFile(project, folderName, fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var tmplFile = "templates/sql/init.tmpl"
	t := template.Must(template.New("init.tmpl").ParseFiles(tmplFile))
	err = t.Execute(file, input)
	if err != nil {
		return err
	}
	return nil
}
