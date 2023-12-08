package generate

import (
	"io"
	"os"
	"strings"
	"text/template"

	"generator/helper"
)

type Interface interface {
	Generate(project, folderName, fileName, tmplPath, tmplFileName string, input any) error
	Append(filePath, tag, append string) error
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

func (r *generate) Append(filePath, tag, append string) error {
	file, err := os.Open(helper.Path() + filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	dataArr := strings.Split(string(data), "\n")

	if err != nil {
		return err
	}

	for i, v := range dataArr {
		if strings.Contains(v, tag) {
			dataArr = helper.InsertIntoSlice(dataArr, i+1, append)
			break
		}
	}
	data = []byte(strings.Join(dataArr, "\n"))

	return os.WriteFile(helper.Path()+filePath, data, 0644)
}
