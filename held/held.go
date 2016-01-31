package held

import (
	"errors"
	"fmt"
	"github.com/Schokomuesl1/bowie/basiswerte"
)

type Held struct {
	Name               string
	Spezies            basiswerte.SpeziesType
	Basiswerte         basiswerte.BerechneteWerte
	AP                 int
	AP_spent           int
	Kultur             basiswerte.KulturType
	Eigenschaften      basiswerte.EigenschaftHandler
	Kampftechniken     basiswerte.KampftechnikHandler
	Talente            basiswerte.TalentHandler
	Sonderfertigkeiten []*basiswerte.Sonderfertigkeit
	Vorteile           []*basiswerte.VorUndNachteil
	Nachteile          []*basiswerte.VorUndNachteil
}

func NewHeld() *Held {
	h := Held{Eigenschaften: *basiswerte.NewEigenschaftHandler(), Kampftechniken: *basiswerte.NewKampftechnikHandler(), Talente: *basiswerte.NewTalentHandler()}
	for k, _ := range basiswerte.AlleEigenschaften {
		h.Eigenschaften.Add(k)
	}
	for _, v := range basiswerte.AlleTalente {
		h.NewTalent(v.Name, v.Probe, v.Steigerungsfaktor, v.Kategorie)
	}
	for _, v := range basiswerte.AlleKampftechniken {
		h.NewKampftechnik(v.Name, v.Typ == "Fernkampf", v.Leiteigenschaft, v.Steigerungsfaktor)
	}
	return &h
}

func (h *Held) SetSpezies(spezies string) error {
	_, existing := basiswerte.AlleSpezies[spezies]
	if !existing {
		return errors.New("Spezies unbekannt!")
	}
	h.Spezies = basiswerte.AlleSpezies[spezies]
	for _, v := range h.Spezies.Vorteile {
		vorteil := basiswerte.GetVorteil(v)
		if vorteil != nil {
			h.Vorteile = append(h.Vorteile, vorteil)
		}
	}

	for _, v := range h.Spezies.Nachteile {
		nachteil := basiswerte.GetNachteil(v)
		if nachteil != nil {
			h.Nachteile = append(h.Nachteile, nachteil)
		}
	}

	h.Basiswerte = *basiswerte.InitBerechneteWerte(&h.Spezies, &h.Eigenschaften)
	return nil
}

func (h *Held) SetKultur(kultur string) error {
	_, existing := basiswerte.AlleKulturen[kultur]
	if !existing {
		return errors.New("Kultur unbekannt!")
	}
	h.Kultur = basiswerte.AlleKulturen[kultur]
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

func (h *Held) APGesamt() int { return (h.AP + h.AP_spent) }

func (h *Held) APAusgeben(menge int) {
	h.AP -= menge
	h.AP_spent += menge
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
	ret += "Abenteuerpunkte\n"
	ret += "-------\n"
	ret += fmt.Sprintf("AP gesamt: %d, AP ausgegeben: %d, AP verfügbar: %d\n", h.AP+h.AP_spent, h.AP_spent, h.AP)
	ret += "Kampftechniken\n"
	ret += "-------\n"
	for _, v := range h.Kampftechniken.Kampftechniken {
		ret += fmt.Sprintf("%s\n", v.String())
	}
	ret += "Kampftechniken\n"
	ret += "-------\n"
	for _, v := range h.Talente.Talente {
		ret += fmt.Sprintf("%s\n", v.String())
	}
	return ret
}

func (h *Held) NewTalent(name string, eigenschaften [3]string, sf string, kat string) bool {
	if h.Talente.Exists(name) {
		return false
	}
	e1 := h.Eigenschaften.Eigenschaften[eigenschaften[0]]
	e2 := h.Eigenschaften.Eigenschaften[eigenschaften[1]]
	e3 := h.Eigenschaften.Eigenschaften[eigenschaften[2]]
	h.Talente.Add(basiswerte.MakeTalent(name, 0, e1, e2, e3, sf, kat))
	return h.Talente.Exists(name)
}

func (h *Held) NewKampftechnik(name string, isFernkampf bool, leiteigenschaften []string, sf string) bool {
	if h.Kampftechniken.Exists(name) {
		return false
	}
	lt := make([]*basiswerte.Eigenschaft, len(leiteigenschaften))
	for i, v := range leiteigenschaften {
		lt[i] = h.Eigenschaften.Eigenschaften[v]
	}
	if isFernkampf {
		h.Kampftechniken.Add(basiswerte.MakeFernkampf(name, 6, lt, h.Eigenschaften.Eigenschaften["FF"], sf))
	} else {
		h.Kampftechniken.Add(basiswerte.MakeNahkampf(name, 6, lt, h.Eigenschaften.Eigenschaften["MU"], sf))
	}
	return h.Kampftechniken.Exists(name)
}
