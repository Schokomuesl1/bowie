package basiswerte

import (
	"encoding/json"
	"io/ioutil"
)

var AlleEigenschaften map[string]string
var AlleSpezies map[string]SpeziesType
var AlleTalente []TalentType
var AlleKulturen map[string]KulturType
var kosten map[string][26]int

func init() {
	file, _ := ioutil.ReadFile("regeln/eigenschaften.json")
	AlleEigenschaften = make(map[string]string)
	json.Unmarshal([]byte(string(file)), &AlleEigenschaften)

	file2, _ := ioutil.ReadFile("regeln/spezies.json")
	speziesTmp := make([]SpeziesType, 0)
	json.Unmarshal([]byte(string(file2)), &speziesTmp)
	AlleSpezies = make(map[string]SpeziesType)
	for _, v := range speziesTmp {
		AlleSpezies[v.Name] = v
	}
	file3, _ := ioutil.ReadFile("regeln/talente.json")
	AlleTalente = make([]TalentType, 0)
	json.Unmarshal([]byte(string(file3)), &AlleTalente)

	file4, _ := ioutil.ReadFile("regeln/kulturen.json")
	kulturTmp := make([]KulturType, 0)
	AlleKulturen = make(map[string]KulturType)
	json.Unmarshal([]byte(string(file4)), &kulturTmp)
	for _, v := range kulturTmp {
		AlleKulturen[v.Name] = v
	}
	file5, _ := ioutil.ReadFile("regeln/kosten.json")
	kostenTmp := make([]SKType, 0)
	kosten = make(map[string][26]int)
	json.Unmarshal([]byte(string(file5)), &kostenTmp)
	for _, v := range kostenTmp {
		kosten[v.SK] = v.Kosten
	}
}

type TalentType struct {
	Name      string
	Kategorie string
	Probe     [3]string
	Belastung string
	Kosten    string
}

type SKType struct {
	SK     string
	Kosten [26]int
}

type EigenschaftenModSpezies struct {
	Eigenschaft []string
	Mod         int
}

type SpeziesType struct {
	Name                       string
	LE                         int
	SK                         int
	ZK                         int
	GS                         int
	EigenschaftsModifikationen []EigenschaftenModSpezies
	Vorteile                   []string
	Nachteile                  []string
	AP                         int
}

type ModPair struct {
	Id    string
	Value int
}

type KulturType struct {
	Name    string
	Talente []ModPair
	Kosten  int
}
