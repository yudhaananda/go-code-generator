package helper

import "os"

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
