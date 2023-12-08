package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/gertd/go-pluralize"
)

func Path() string {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	dirnameArr := strings.Split(dirname, "/")

	path := ""
	for i := 0; i < len(dirnameArr)-2; i++ {
		path += dirnameArr[i] + "/"
	}
	return path
}

func InsertIntoSlice(slice []string, index int, value string) []string {
	// Ensure the index is within bounds
	if index < 0 || index > len(slice) {
		return slice
	}

	// Create a new slice with enough capacity for the new element
	newSlice := make([]string, len(slice)+1)

	// Copy the elements before the insertion point
	copy(newSlice[:index], slice[:index])

	// Insert the new element
	newSlice[index] = value

	// Copy the elements after the insertion point
	copy(newSlice[index+1:], slice[index:])

	return newSlice
}

func CreateAdditionalFile(project, folderName, fileName string) (*os.File, error) {
	err := os.MkdirAll(Path()+"/"+folderName, os.ModePerm)
	if err != nil {
		return &os.File{}, err
	}
	file, err := os.Create(Path() + "/" + folderName + "/" + fileName)

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

func ConvertToRepositoriesInjection(input string) string {
	return fmt.Sprintf("		%s: %s.Init(%s.Param{Db: param.Db, TableName: \"%s\"}),",
		input,
		ConvertToLowerCase(input),
		ConvertToLowerCase(input),
		ConvertToSnakeCase(ConvertToPlural(input)),
	)
}

func ConvertToRepositoriesInjectionImport(input string, projectName string) string {
	return fmt.Sprintf("	%s \"%s/src/repositories/%s\"",
		ConvertToLowerCase(input),
		projectName,
		ConvertToSnakeCase(input),
	)
}

func ConvertToRepositoriesInjectionStruct(input string) string {
	return fmt.Sprintf("	%s %s.Interface",
		input,
		ConvertToLowerCase(input),
	)
}

func ConvertToServicesInjection(input string) string {
	return fmt.Sprintf(
		"		%s: %s.Init(%s.Param{%sRepository: param.Repositories.%s}),",
		input,
		ConvertToLowerCase(input),
		ConvertToLowerCase(input),
		input,
		input,
	)
}

func ConvertToServicesInjectionStruct(input string) string {
	return fmt.Sprintf("	%s %s.Interface",
		input,
		ConvertToLowerCase(input),
	)
}

func ConvertToServicesInjectionImport(input string, projectName string) string {
	return fmt.Sprintf("	%s \"%s/src/services/%s\"",
		ConvertToLowerCase(input),
		projectName,
		ConvertToSnakeCase(input),
	)
}

func ConvertToInitHandler(input string) string {
	return fmt.Sprintf("	%sApi := api.Group(\"/%s\").Use(h.middleware.AuthMiddleware)\n	{\n		%sApi.GET(\"/\", h.Get%s)\n		%sApi.POST(\"/\", h.Create%s)\n		%sApi.PUT(\"/:id\", h.Update%s)\n		%sApi.DELETE(\"/:id\", h.Delete%s)\n	}",
		ConvertToCamelCase(input),
		ConvertToDash(input),
		ConvertToCamelCase(input),
		input,
		ConvertToCamelCase(input),
		input,
		ConvertToCamelCase(input),
		input,
		ConvertToCamelCase(input),
		input,
	)
}
