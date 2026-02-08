package file

import (
	"errors"
	"os"
	"path/filepath"
)

//Чтение любого файла
//Проверка что это json расширение файла

func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if filepath.Ext(filename) != ".json" {
		return "", errors.New("invalid file extension")
	}
	return string(content), nil
}
