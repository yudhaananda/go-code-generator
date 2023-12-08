package zipping

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/yudhaananda/go-code-generator/helper"
)

type Interface interface {
	Delete(projectName string) error
	Zipping(projectName string) ([]byte, error)
}

type zipping struct {
}

func Init() Interface {
	return &zipping{}
}

func (s *zipping) Delete(projectName string) error {
	err := os.RemoveAll(projectName)
	if err != nil {
		return err
	}
	err = os.Remove(projectName + ".zip")
	if err != nil {
		return err
	}
	return nil
}

func (s *zipping) Zipping(projectName string) ([]byte, error) {
	baseFolder := projectName + "/"

	zipName, err := os.Create(projectName + ".zip")

	if err != nil {
		return nil, err
	}

	defer zipName.Close()

	w := zip.NewWriter(zipName)

	s.addFiles(w, baseFolder)

	err = w.Close()
	if err != nil {
		return nil, err
	}
	zipFile, err := os.ReadFile(projectName + ".zip")

	if err != nil {
		return nil, err
	}

	return zipFile, nil
}

func (s *zipping) addFiles(w *zip.Writer, basePath string) {
	// Open the Directory
	files, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		// fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := os.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			fileName := file.Name()
			if strings.Contains(file.Name(), ".txt") {
				split := strings.Split(file.Name(), ".")
				fileName = strings.Join(helper.Remove(split, len(split)-1), ".")
			}

			// Add some files to the archive.
			var f io.Writer
			f, err = w.Create(basePath + fileName)

			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(basePath + file.Name())
		} else if file.IsDir() {

			newBase := basePath + file.Name() + "/"

			s.addFiles(w, newBase)
		}
	}
}
