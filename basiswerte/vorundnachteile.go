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
