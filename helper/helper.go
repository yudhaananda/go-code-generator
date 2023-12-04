package helper

import (
	"os"
	"regexp"
	"strings"

	"github.com/gertd/go-pluralize"
)

func CreateFile(project, folderName, fileName string) (*os.File, error) {
	err := os.MkdirAll(project+"/src/"+folderName, os.ModePerm)
	if err != nil {
		return &os.File{}, err
	}
	file, err := os.Create(project + "/src/" + folderName + "/" + fileName)

	if err != nil {
		return file, err
	}
	return file, nil
}

func CreateAdditionalFile(project, folderName, fileName string) (*os.File, error) {
	err := os.MkdirAll(project+"/"+folderName, os.ModePerm)
	if err != nil {
		return &os.File{}, err
	}
	file, err := os.Create(project + "/" + folderName + "/" + fileName)

	if err != nil {
		return file, err
	}
	return file, nil
}

func ConvertToCamelCase(input string) string {
	split := strings.Split(input, "")
	split[0] = strings.ToLower(split[0])
	return strings.Join(split, "")
}

func ConvertToSnakeCase(input string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	submatchall := re.FindAllString(input, -1)
	return strings.ToLower(strings.Join(submatchall, "_"))
}

func ConvertToDash(input string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	submatchall := re.FindAllString(input, -1)
	return strings.ToLower(strings.Join(submatchall, "-"))
}

func ConvertToLowerCase(input string) string {
	return strings.ToLower(input)
}

func ConvertToPlural(input string) string {
	plural := pluralize.NewClient()
	return plural.Plural(input)
}
