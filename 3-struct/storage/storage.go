package storage

import (
	"errors"
	"os"
	"path/filepath"
)

//Сохранение bin в виде json в локальном файле
//Чтение списка bin в виде json из локального файла

type Storage interface {
	SaveFile(filename string, data []byte) error
	ReadFile(filename string) ([]byte, error)
}

type LocalStorage struct{}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (s *LocalStorage) SaveFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

func (s *LocalStorage) ReadFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if filepath.Ext(filename) != ".json" {
		return nil, errors.New("invalid file extension")
	}
	return content, nil
}

// func SaveBin(bin interface{}, filename string) error {
// 	jsonData, err := json.Marshal(bin)
// 	if err != nil {
// 		return err
// 	}
// 	return os.WriteFile(filename, jsonData, 0644)
// }

// func ReadBins(filename string) ([]byte, error) {
// 	content, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if filepath.Ext(filename) != ".json" {
// 		return nil, errors.New("invalid file extension")
// 	}
// 	return content, nil
// }
