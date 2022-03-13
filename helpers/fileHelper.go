package helpers

import (
	"errors"
	"os"
)

func MakeFolderIfNotExists(folderPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0755)
	}
}

func MakeFileIfNotExists(filePath string) (*os.File, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return os.Create(filePath)
	}
	return nil, errors.New("file already exists")
}

func WriteFile(filePath string, content string) error {
	file, err := MakeFileIfNotExists(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}
