package bins

import (
	"encoding/json"
	"homework/3-struct/storage"
)

type BinService struct {
	storage storage.Storage // DI через интерфейс!
}

func NewBinService(s storage.Storage) *BinService {
	return &BinService{storage: s}
}

func (bs *BinService) SaveBin(bin interface{}, filename string) error {
	jsonData, err := json.Marshal(bin)
	if err != nil {
		return err
	}
	return bs.storage.SaveFile(filename, jsonData)
}

func (bs *BinService) ReadBins(filename string) ([]byte, error) {
	return bs.storage.ReadFile(filename)
}
