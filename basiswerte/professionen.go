package basiswerte

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var AlleProfessionen []Profession

type Profession struct {
	Name                  string
	Kommentar             string
	Voraussetzungen       VoraussetzungenProfession
	Kampftechniken        [][2]string
	KampftechnikenAuswahl []KampftechnikWahlPack
	Talente               [][2]string
	Sonderfertigkeiten    []string
	APKosten              int
}
type KampftechnikWahlPack struct {
	Wahlmoeglichkeiten []string
	ZuWaehlen          int
	Wert               int
}
type VoraussetzungenProfession struct {
	Eigenschaften [][2]string
	Kultur        []string
}

func init() {
	AlleProfessionen = make([]Profession, 0)
	readAndAddToProfessionen("regeln/professionen/weltliche.json")
	fmt.Println(AlleProfessionen)
}

func readAndAddToProfessionen(filename string) {
	file, _ := ioutil.ReadFile(filename)
	var tmpList = make([]Profession, 0)
	json.Unmarshal([]byte(string(file)), &tmpList)
	AlleProfessionen = append(AlleProfessionen, tmpList...)
}
