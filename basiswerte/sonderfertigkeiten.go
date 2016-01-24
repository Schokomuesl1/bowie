package basiswerte

import (
	"encoding/json"
	"io/ioutil"
)

var AllgemeineSF []Sonderfertigkeit

type Sonderfertigkeit struct {
	Name             string
	Vorraussetzungen VorraussetzungSF
	APKosten         int
}

type VorraussetzungSF struct {
	Talente          [][2]string
	Vorteile         []string
	NichtNachteil    []string
	Eigenschaften    [][2]string
	Sonderfertigkeit []string
	Sonstiges        string
}

func init() {
	vorteileFile, _ := ioutil.ReadFile("regeln/sonderfertigkeiten/allgemeine.json")
	AllgemeineSF = make([]Sonderfertigkeit, 0)
	json.Unmarshal([]byte(string(vorteileFile)), &AllgemeineSF)
}
