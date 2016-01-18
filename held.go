package bowie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var AlleEigenschaften map[string]string

func init() {
	file, _ := ioutil.ReadFile("data/eigenschaften.json")
	AlleEigenschaften = make(map[string]string)
	json.Unmarshal([]byte(string(file)), &AlleEigenschaften)
}

type modPair struct {
	Id    string
	Value int32
}
type KulturType struct {
	Name          string
	Modifications []modPair
}

type EigenschaftHandler struct {
	Eigenschaften map[string]*Eigenschaft
}

func (e *EigenschaftHandler) Set(name string, value int) bool {
	if !e.exists(name) {
		e.Add(name)
	}
	e.Eigenschaften[name].Set(value)
	return true
}

func (e *EigenschaftHandler) exists(name string) bool {
	_, existing := e.Eigenschaften[name]
	return existing
}

func NewEigenschaftHandler() *EigenschaftHandler {
	return &EigenschaftHandler{Eigenschaften: make(map[string]*Eigenschaft)}
}

func (e *EigenschaftHandler) Add(name string) bool {
	if e.exists(name) {
		return false
	}
	e.Eigenschaften[name] = MakeEigenschaft(name, 0)
	return true
}

type TalentHandler struct {
	Talente map[string]*Talent
}

func NewTalentHandler() *TalentHandler {
	return &TalentHandler{Talente: make(map[string]*Talent)}
}

func (t *TalentHandler) Exists(name string) bool {
	_, existing := t.Talente[name]
	return existing
}

func (t *TalentHandler) Add(talent *Talent) bool {
	if t.Exists(talent.Name) {
		return false
	}
	t.Talente[talent.Name] = talent
	return true
}

type Held struct {
	Name          string
	Spezies       string
	Kultur        KulturType
	Eigenschaften EigenschaftHandler
	Talente       TalentHandler
}

func NewHeld() *Held {
	h := Held{Eigenschaften: *NewEigenschaftHandler(), Talente: *NewTalentHandler()}
	for k, _ := range AlleEigenschaften {
		h.Eigenschaften.Add(k)
	}
	return &h
}

//String prints an overview of the hero
func (h *Held) String() string {
	ret := fmt.Sprintf("Name: %s, Kultur: %s\n", h.Name, h.Kultur.Name)
	ret += "Eigenschaften\n"
	ret += "-------------\n"
	for _, v := range h.Eigenschaften.Eigenschaften {
		ret += fmt.Sprintf("%s: %d\n", v.Name, v.Value())
	}
	ret += "Talente\n"
	ret += "-------\n"
	for _, v := range h.Talente.Talente {
		ret += fmt.Sprintf("%s\n", v.String())
	}
	return ret
}

func (h *Held) NewTalent(name string, eigenschaften [3]string) bool {
	if h.Talente.Exists(name) {
		return false
	}
	e1 := h.Eigenschaften.Eigenschaften[eigenschaften[0]]
	e2 := h.Eigenschaften.Eigenschaften[eigenschaften[1]]
	e3 := h.Eigenschaften.Eigenschaften[eigenschaften[2]]
	h.Talente.Add(MakeTalent(name, 0, e1, e2, e3))
	if h.Talente.Exists(name) {
		return true
	}
	return false
}

type TalentParse struct {
	Name      string
	Kategorie string
	Probe     [3]string
	Belastung string
	Kosten    string
}

type EigenschaftenModSpezies struct {
	Eigenschaft []string
	Mod         int
}

type Spezies struct {
	Name                       string
	LE                         int
	SK                         int
	GS                         int
	EigenschaftsModifikationen []EigenschaftenModSpezies
	Vorteile                   []string
	Nachteile                  []string
	AP                         int
}
