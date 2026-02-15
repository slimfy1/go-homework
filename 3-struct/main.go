package main

import (
	"fmt"
	"homework/3-struct/api"
	"homework/3-struct/bins"
	"homework/3-struct/config"
	"homework/3-struct/storage"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAT time.Time
	Name      string
}

type BinList struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func main() {
	configEnv := config.NewConfig()
	apiClient := api.NewClient(configEnv)
	fmt.Println("configEnv:", apiClient)
	// 1. Создаём реализацию Storage
	localStorage := storage.NewLocalStorage()

	// 2. Внедряем зависимость через интерфейс
	binService := bins.NewBinService(localStorage)

	// 3. Используем сервис
	testBin := createBinStruct("123", false, "Test Bin")

	err := binService.SaveBin(testBin, "data.json")
	if err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	data, err := binService.ReadBins("data.json")
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Success! Data:", string(data))
}

func createBinStruct(id string, private bool, name string) Bin {

	binList := Bin{
		Id:        id,
		Private:   private,
		CreatedAT: time.Now(),
		Name:      name,
	}
	return binList

}

func createBinListStruct(id string, private bool, name string) BinList {
	return BinList{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
