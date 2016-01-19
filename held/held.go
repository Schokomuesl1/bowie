package bowie

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
)

var AlleEigenschaften map[string]string
var AlleSpezies map[string]SpeziesType
var AlleTalente []TalentType
var AlleKulturen map[string]KulturType

func init() {
    file, _ := ioutil.ReadFile("data/eigenschaften.json")
    AlleEigenschaften = make(map[string]string)
    json.Unmarshal([]byte(string(file)), &AlleEigenschaften)

    file2, _ := ioutil.ReadFile("data/spezies.json")
    speziesTmp := make([]SpeziesType, 0)
    json.Unmarshal([]byte(string(file2)), &speziesTmp)
    AlleSpezies = make(map[string]SpeziesType)
    for _, v := range speziesTmp {
        AlleSpezies[v.Name] = v
    }
    file3, _ := ioutil.ReadFile("data/talente.json")
    AlleTalente = make([]TalentType, 0)
    json.Unmarshal([]byte(string(file3)), &AlleTalente)
    file4, _ := ioutil.ReadFile("data/kulturen.json")
    kulturTmp := make([]KulturType, 0)
    AlleKulturen = make(map[string]KulturType)
    json.Unmarshal([]byte(string(file4)), &kulturTmp)
    for _, v := range kulturTmp {
        AlleKulturen[v.Name] = v
    }
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

type EigenschaftHandler struct {
    Eigenschaften map[string]*Eigenschaft
}

func (e *EigenschaftHandler) Set(name string, value int) bool {
    if !e.Exists(name) {
        e.Add(name)
    }
    e.Eigenschaften[name].Set(value)
    return true
}

func (e *EigenschaftHandler) Exists(name string) bool {
    _, existing := e.Eigenschaften[name]
    return existing
}

func NewEigenschaftHandler() *EigenschaftHandler {
    return &EigenschaftHandler{Eigenschaften: make(map[string]*Eigenschaft)}
}

func (e *EigenschaftHandler) Add(name string) bool {
    if e.Exists(name) {
        return false
    }
    e.Eigenschaften[name] = MakeEigenschaft(name, 0)
    return true
}

func (e *EigenschaftHandler) Get(name string) *Eigenschaft {
    if !e.Exists(name) {
        return nil
    }
    return e.Eigenschaften[name]
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

func (t *TalentHandler) Get(talent string) *Talent {
    if !t.Exists(talent) {
        return nil
    }
    return t.Talente[talent]
}

//func MakeCalculatedDependentValue(name string, mult float32, base []DependentValue) *CalculatedDependentValue {

type Basiswert struct {
    Grundwert     int
    Berechnung    *CalculatedDependentValue
    Modifikatoren []int // werden gesamt addiert/subtrahiert
}

func MakeBasiswert(name string, mult float32, base []DependentValue, grundwert int) *Basiswert {
    return &Basiswert{Grundwert: grundwert, Berechnung: MakeCalculatedDependentValue(name, mult, base), Modifikatoren: make([]int, 0)}
}

func (b *Basiswert) Value() int {
    tmp := b.Grundwert + b.Berechnung.Value()
    for _, v := range b.Modifikatoren {
        tmp += v
    }
    return tmp
}

func (b *Basiswert) AddModifier(value int) {
    n := len(b.Modifikatoren)
    if n == cap(b.Modifikatoren) {
        newMod := make([]int, len(b.Modifikatoren), 2*len(b.Modifikatoren)+1)
        copy(newMod, b.Modifikatoren)
        b.Modifikatoren = newMod
    }
    b.Modifikatoren = b.Modifikatoren[0 : n+1]
    b.Modifikatoren[n] = value
}

type BerechneteWerte struct {
    Lebensenergie   *Basiswert // Grundwert + 2*KO + Vor/Nachteile
    Astralenergie   *Basiswert // Grundwert + Leiteigenschaft + Vor/Nachteile
    Karmaenergie    *Basiswert // Grundwert + Leiteigenschaft + Vor/Nachteile
    Seelenkraft     *Basiswert // Grundwert + (MU + KL + IN)/6 + Vor/Nachteile
    Zaehigkeit      *Basiswert // Grundwert + (KO + KO + KK)/6 + Vor/Nachteile
    Ausweichen      *Basiswert // GE / 2
    Initiative      *Basiswert // (MU + GE)/2 + Vor/Nachteile
    Geschwindigkeit *Basiswert // Grundwert + Vor/Nachteile)
}

func InitBerechneteWerte(spezies *SpeziesType, eigenschaften *EigenschaftHandler) *BerechneteWerte {
    return &BerechneteWerte{
        Lebensenergie:   MakeBasiswert("Lebensenergie", 2, []DependentValue{eigenschaften.Get("KO")}, spezies.LE),
        Astralenergie:   nil, // check this
        Karmaenergie:    nil, // check this
        Seelenkraft:     MakeBasiswert("Seelenkraft", (1.0 / 6.0), []DependentValue{eigenschaften.Get("MU"), eigenschaften.Get("KL"), eigenschaften.Get("IN")}, spezies.SK),
        Zaehigkeit:      MakeBasiswert("Zaehigkeit", (1.0 / 6.0), []DependentValue{eigenschaften.Get("KO"), eigenschaften.Get("KO"), eigenschaften.Get("KK")}, spezies.ZK),
        Ausweichen:      MakeBasiswert("Ausweichen", 0.5, []DependentValue{eigenschaften.Get("GE")}, 0),
        Initiative:      MakeBasiswert("Initiative", 0.5, []DependentValue{eigenschaften.Get("GE"), eigenschaften.Get("MU")}, 0),
        Geschwindigkeit: MakeBasiswert("Geschwindigkeit", 1, []DependentValue{}, spezies.GS)}
}

type Held struct {
    Name          string
    Spezies       SpeziesType
    Basiswerte    BerechneteWerte
    Kultur        KulturType
    Eigenschaften EigenschaftHandler
    Talente       TalentHandler
}

func NewHeld() *Held {
    h := Held{Eigenschaften: *NewEigenschaftHandler(), Talente: *NewTalentHandler()}
    for k, _ := range AlleEigenschaften {
        h.Eigenschaften.Add(k)
    }
    for _, v := range AlleTalente {
        h.NewTalent(v.Name, v.Probe)
    }
    return &h
}

func (h *Held) SetSpezies(spezies string) error {
    _, existing := AlleSpezies[spezies]
    if !existing {
        return errors.New("Spezies unbekannt!")
    }
    h.Spezies = AlleSpezies[spezies]
    h.Basiswerte = *InitBerechneteWerte(&h.Spezies, &h.Eigenschaften)
    return nil
}

func (h *Held) SetKultur(kultur string) error {
    _, existing := AlleKulturen[kultur]
    if !existing {
        return errors.New("Kultur unbekannt!")
    }
    h.Kultur = AlleKulturen[kultur]
    for _, v := range h.Kultur.Talente {
        tmp := h.Talente.Get(v.Id)
        if tmp == nil {
            fmt.Printf("Talent %s not found!\n", v.Id)
        } else {
            tmp.AddValue(v.Value)
        }

    }
    return nil
}

//String prints an overview of the hero
func (h *Held) String() string {
    ret := fmt.Sprintf("Name: %s, Kultur: %s\n", h.Name, h.Kultur.Name)
    ret += "Eigenschaften\n"
    ret += "-------------\n"
    for _, v := range h.Eigenschaften.Eigenschaften {
        ret += fmt.Sprintf("%s: %d\n", v.Name, v.Value())
    }
    ret += "Basiswerte\n"
    ret += "----------\n"
    if h.Basiswerte.Lebensenergie != nil {
        ret += fmt.Sprintf("Lebensenergie: %d\n", h.Basiswerte.Lebensenergie.Value())
    } else {
        ret += "Lebensenergie: -\n"
    }
    if h.Basiswerte.Astralenergie != nil {
        ret += fmt.Sprintf("Astralenergie: %d\n", h.Basiswerte.Astralenergie.Value())
    } else {
        ret += "Astralenergie: -\n"
    }
    if h.Basiswerte.Karmaenergie != nil {
        ret += fmt.Sprintf("Karmaenergie: %d\n", h.Basiswerte.Karmaenergie.Value())
    } else {
        ret += "Karmaenergie: -\n"
    }
    if h.Basiswerte.Seelenkraft != nil {
        ret += fmt.Sprintf("Seelenkraft: %d\n", h.Basiswerte.Seelenkraft.Value())
    } else {
        ret += "Seelenkraft: -\n"
    }
    if h.Basiswerte.Zaehigkeit != nil {
        ret += fmt.Sprintf("Zaehigkeit: %d\n", h.Basiswerte.Zaehigkeit.Value())
    } else {
        ret += "Zaehigkeit: -\n"
    }
    if h.Basiswerte.Ausweichen != nil {
        ret += fmt.Sprintf("Ausweichen: %d\n", h.Basiswerte.Ausweichen.Value())
    } else {
        ret += "Ausweichen: -\n"
    }
    if h.Basiswerte.Initiative != nil {
        ret += fmt.Sprintf("Initiative: %d\n", h.Basiswerte.Initiative.Value())
    } else {
        ret += "Initiative: -\n"
    }
    if h.Basiswerte.Geschwindigkeit != nil {
        ret += fmt.Sprintf("Geschwindigkeit: %d\n", h.Basiswerte.Geschwindigkeit.Value())
    } else {
        ret += "Geschwindigkeit: -\n"
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

type TalentType struct {
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
