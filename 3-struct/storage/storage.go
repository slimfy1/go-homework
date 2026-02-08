package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

//Сохранение bin в виде json в локальном файле
//Чтение списка bin в виде json из локального файла

func saveBin(bin []byte, filename string) error {
	json, err := json.Marshal(bin)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

func readBins(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if filepath.Ext(filename) != ".json" {
		return nil, errors.New("invalid file extension")
	}
	return content, nil
}
