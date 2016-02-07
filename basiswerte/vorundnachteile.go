package basiswerte

import (
	"encoding/json"
	"io/ioutil"
)

var Vorteile []VorUndNachteil
var Nachteile []VorUndNachteil

type VorUndNachteil struct {
	Name             string
	Vorraussetzungen VorraussetzungVorteilNachteil
	APKosten         int
	FromSpezies      bool // initializes by default to false which is good :P
}

type VorraussetzungVorteilNachteil struct {
	Vorteile       []string
	Nachteile      []string
	NichtVorteile  []string
	NichtNachteile []string
	Sonstiges      string
}

func init() {
	vorteileFile, _ := ioutil.ReadFile("regeln/VorUndNachteile/vorteile.json")
	Vorteile = make([]VorUndNachteil, 0)
	json.Unmarshal([]byte(string(vorteileFile)), &Vorteile)

	nachteileFile, _ := ioutil.ReadFile("regeln/VorUndNachteile/nachteile.json")
	Nachteile = make([]VorUndNachteil, 0)
	json.Unmarshal([]byte(string(nachteileFile)), &Nachteile)
}

func (vn *VorUndNachteil) DeleteButtonIfApplicable() string {
	if !vn.FromSpezies {
		return "text-danger glyphicon glyphicon-remove"
	}
	return ""
}

func GetVorteil(name string) *VorUndNachteil {
	for _, v := range Vorteile {
		if name == v.Name {
			return &v
		}
	}
	return nil
}

func GetNachteil(name string) *VorUndNachteil {
	for _, v := range Nachteile {
		if name == v.Name {
			return &v
		}
	}
	return nil
}
