package main

import "time"

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
