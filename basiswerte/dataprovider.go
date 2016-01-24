package basiswerte

import (
	"encoding/json"
	"io/ioutil"
)

var AlleEigenschaften map[string]string
var AlleSpezies map[string]SpeziesType
var AlleTalente []TalentType
var AlleKulturen map[string]KulturType
var AlleLiturgien map[string]LiturgieType
var AlleKampftechniken map[string]KampftechnikType
var Kostentable map[string][26]int

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
	Kostentable = make(map[string][26]int)
	json.Unmarshal([]byte(string(file5)), &kostenTmp)
	for _, v := range kostenTmp {
		Kostentable[v.SK] = v.APKosten
	}

	file6, _ := ioutil.ReadFile("regeln/liturgien.json")
	liturgieTmp := make([]LiturgieType, 0)
	AlleLiturgien = make(map[string]LiturgieType)
	json.Unmarshal([]byte(string(file6)), &liturgieTmp)
	for _, v := range liturgieTmp {
		AlleLiturgien[v.Name] = v
	}
	file7, _ := ioutil.ReadFile("regeln/kampftechniken.json")
	kampftechnikTmp := make([]KampftechnikType, 0)
	AlleKampftechniken = make(map[string]KampftechnikType)
	json.Unmarshal([]byte(string(file7)), &kampftechnikTmp)
	for _, v := range kampftechnikTmp {
		AlleKampftechniken[v.Name] = v
	}

}

type TalentType struct {
	Name              string
	Kategorie         string
	Probe             [3]string
	Belastung         string
	Steigerungsfaktor string
}

type SKType struct {
	SK       string
	APKosten [26]int
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
	APKosten                   int
}

type ModPair struct {
	Id    string
	Value int
}

type KulturType struct {
	Name     string
	Talente  []ModPair
	APKosten int
}

type LiturgieType struct {
	Name              string
	Probe             ProbeMitMod
	Wirkung           string
	Dauer             string
	Kosten            []string
	Reichweite        ReichweiteMitMod
	Wirkungsdauer     string
	Zielkategorie     []string
	Verbreitung       [][2]string
	Steigerungsfaktor string
}

type ProbeMitMod struct {
	Eigenschaften [3]string
	Mod           string
}

type ReichweiteMitMod struct {
	Reichweite string
	Mod        string
}
