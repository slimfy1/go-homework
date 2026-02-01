package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAT time.Time
	name      string
}

type BinList struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func main() {

}

func createBinStruct(id string, private bool, name string) Bin {

	binList := Bin{
		id:        id,
		private:   private,
		createdAT: time.Now(),
		name:      name,
	}
	return binList

}

func createBinListStruct(id string, private bool, name string) BinList {
	return BinList{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}
