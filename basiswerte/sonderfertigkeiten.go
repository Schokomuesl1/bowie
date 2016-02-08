package basiswerte

import (
	"encoding/json"
	"io/ioutil"
)

var AllgemeineSF []Sonderfertigkeit
var KarmaleSF []Sonderfertigkeit
var MagischeSF []Sonderfertigkeit
var KampfSF []Sonderfertigkeit
var SprachenSF []Sonderfertigkeit
var SchriftenSF []Sonderfertigkeit

type SFType int

const (
	UNBEKANNT SFType = iota
	ALLGEMEIN
	KARMAL
	MAGISCH
	KAMPF
	SPRACHE
	SCHRIFT
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
	readAndAppendToSFList("regeln/sonderfertigkeiten/fertigkeitsspezialisierungen.json", &AllgemeineSF)
	readAndMakeSFList("regeln/sonderfertigkeiten/sprachen.json", &SprachenSF)
	readAndMakeSFList("regeln/sonderfertigkeiten/schriften.json", &SchriftenSF)
}

func readAndMakeSFList(filename string, sfList *[]Sonderfertigkeit) {
	file, _ := ioutil.ReadFile(filename)
	*sfList = make([]Sonderfertigkeit, 0)
	json.Unmarshal([]byte(string(file)), sfList)
}

func readAndAppendToSFList(filename string, sfList *[]Sonderfertigkeit) {
	file, _ := ioutil.ReadFile(filename)
	tmp := make([]Sonderfertigkeit, 0)
	json.Unmarshal([]byte(string(file)), &tmp)
	*sfList = append(*sfList, tmp...)
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
	for _, v := range SprachenSF {
		if name == v.Name {
			return SPRACHE
		}
	}
	for _, v := range SchriftenSF {
		if name == v.Name {
			return SCHRIFT
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
	for _, v := range SprachenSF {
		if name == v.Name {
			return &v
		}
	}
	for _, v := range SchriftenSF {
		if name == v.Name {
			return &v
		}
	}
	// more groups here...
	return nil
}
