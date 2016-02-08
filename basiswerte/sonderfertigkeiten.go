package basiswerte

import (
	"encoding/json"
	"io/ioutil"
)

var AllgemeineSF []Sonderfertigkeit
var KarmaleSF []Sonderfertigkeit
var MagischeSF []Sonderfertigkeit
var KampfSF []Sonderfertigkeit

type SFType int

const (
	UNBEKANNT SFType = iota
	ALLGEMEIN
	KARMAL
	MAGISCH
	KAMPF
)

type Sonderfertigkeit struct {
	Name             string
	Vorraussetzungen VorraussetzungSF
	Kampftechnik     []string
	APKosten         int
}

type VorraussetzungSF struct {
	Talente            [][2]string
	Vorteile           []string
	NichtNachteil      []string
	Eigenschaften      [][2]string
	Sonderfertigkeiten []string
	Sonstiges          string
}

func init() {
	readAndMakeSFList("regeln/sonderfertigkeiten/allgemeine.json", &AllgemeineSF)
	readAndMakeSFList("regeln/sonderfertigkeiten/karmale.json", &KarmaleSF)
	readAndMakeSFList("regeln/sonderfertigkeiten/magische.json", &MagischeSF)
	readAndMakeSFList("regeln/sonderfertigkeiten/kampf.json", &KampfSF)
}

func readAndMakeSFList(filename string, sfList *[]Sonderfertigkeit) {
	file, _ := ioutil.ReadFile(filename)
	*sfList = make([]Sonderfertigkeit, 0)
	json.Unmarshal([]byte(string(file)), sfList)
}

func GetSFType(name string) SFType {
	for _, v := range AllgemeineSF {
		if name == v.Name {
			return ALLGEMEIN
		}
	}
	for _, v := range KarmaleSF {
		if name == v.Name {
			return KARMAL
		}
	}
	for _, v := range MagischeSF {
		if name == v.Name {
			return MAGISCH
		}
	}
	for _, v := range KampfSF {
		if name == v.Name {
			return KAMPF
		}
	}
	return UNBEKANNT
}

func GetSF(name string) *Sonderfertigkeit {
	for _, v := range AllgemeineSF {
		if name == v.Name {
			return &v
		}
	}
	for _, v := range KarmaleSF {
		if name == v.Name {
			return &v
		}
	}
	for _, v := range MagischeSF {
		if name == v.Name {
			return &v
		}
	}
	for _, v := range KampfSF {
		if name == v.Name {
			return &v
		}
	}
	// more groups here...
	return nil
}
