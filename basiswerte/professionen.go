package basiswerte

import (
	"encoding/json"
	"io/ioutil"
	"sort"
)

var AlleProfessionen ProfessionsListe

type Profession struct {
	Name                  string
	Kommentar             string
	Voraussetzungen       VoraussetzungenProfession
	Kampftechniken        [][2]string
	KampftechnikenAuswahl []WahlPack
	Zauber                [][3]string
	ZauberAuswahl         []WahlPack
	Liturgien             [][2]string
	LiturgienAuswahl      []WahlPack
	Talente               [][2]string
	Sonderfertigkeiten    []string
	APKosten              int
}

var DummyProfession Profession

type WahlPack struct {
	Wahlmoeglichkeiten []string
	ZuWaehlen          int
	Wert               int
}

type VoraussetzungenProfession struct {
	Eigenschaften      [][2]string
	Kultur             []string
	Sonderfertigkeiten []string
	Vorteil            []string
	Spezies            []string
}

func init() {
	DummyProfession = Profession{Name: "-"}
}

func (vp *VoraussetzungenProfession) KulturOK(kultur string) bool {
	if len(vp.Kultur) == 0 {
		return true
	}
	for _, v := range vp.Kultur {
		if kultur == v {
			return true
		}
	}
	return false
}

func (vp *VoraussetzungenProfession) SpeziesOK(spezies string) bool {
	if len(vp.Spezies) == 0 {
		return true
	}
	for _, v := range vp.Spezies {
		if spezies == v {
			return true
		}
	}
	return false
}

func init() {
	AlleProfessionen = make([]Profession, 0)
	readAndAddToProfessionen("regeln/professionen/weltliche.json")
	readAndAddToProfessionen("regeln/professionen/magische.json")
}

func readAndAddToProfessionen(filename string) {
	file, _ := ioutil.ReadFile(filename)
	var tmpList = make([]Profession, 0)
	json.Unmarshal([]byte(string(file)), &tmpList)
	AlleProfessionen = append(AlleProfessionen, tmpList...)
}

type ProfessionsListe []Profession

type ProfessionsListePtr []*Profession

func (p ProfessionsListePtr) Len() int           { return len(p) }
func (p ProfessionsListePtr) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ProfessionsListePtr) Less(i, j int) bool { return p[i].Name < p[j].Name }

func (p ProfessionsListe) NachKulturUndSpezies(kultur string, spezies string) (pl ProfessionsListePtr) {
	pl = make(ProfessionsListePtr, 0)
	for i := 0; i < len(p); i++ {
		if p[i].Voraussetzungen.KulturOK(kultur) && p[i].Voraussetzungen.SpeziesOK(spezies) {
			pl = append(pl, &p[i])
		}
	}
	sort.Sort(pl)
	return pl
}

func (p *ProfessionsListePtr) NachName(name string) *Profession {
	for _, v := range *p {
		if v.Name == name {
			return v
		}
	}
	return nil
}
