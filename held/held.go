package held

import (
	"errors"
	"fmt"
	"github.com/Schokomuesl1/bowie/basiswerte"
	"strconv"
)

type SFListe struct {
	Allgemeine []*basiswerte.Sonderfertigkeit
	Karmale    []*basiswerte.Sonderfertigkeit
	Magische   []*basiswerte.Sonderfertigkeit
	Kampf      []*basiswerte.Sonderfertigkeit
	Schriften  []*basiswerte.Sonderfertigkeit
	Sprachen   []*basiswerte.Sonderfertigkeit
}
type Held struct {
	Name               string
	Spezies            basiswerte.SpeziesType
	Basiswerte         basiswerte.BerechneteWerte
	AP                 int
	AP_spent           int
	Kultur             basiswerte.KulturType
	Profession         basiswerte.Profession
	Eigenschaften      basiswerte.EigenschaftHandler
	Kampftechniken     basiswerte.KampftechnikHandler
	Talente            basiswerte.TalentHandler
	Sonderfertigkeiten SFListe
	Vorteile           []basiswerte.VorUndNachteil
	Nachteile          []basiswerte.VorUndNachteil
	Liturgien          basiswerte.LiturgieHandler
	Zauber             basiswerte.ZauberHandler
	ZauberCount        [2]int
	ZauberCountMax     [2]int
	Erfahrungsgrad     basiswerte.Erfahrungsgrad
}

func NewHeld() *Held {
	h := Held{Eigenschaften: *basiswerte.NewEigenschaftHandler(), Kampftechniken: *basiswerte.NewKampftechnikHandler(), Talente: *basiswerte.NewTalentHandler(), Zauber: *basiswerte.NewZauberHandler(), Liturgien: *basiswerte.NewLiturgieHandler(), Profession: basiswerte.DummyProfession}
	for k, _ := range basiswerte.AlleEigenschaften {
		h.Eigenschaften.Add(k)
	}
	for _, v := range basiswerte.AlleTalente {
		h.NewTalent(v.Name, v.Probe, v.Steigerungsfaktor, v.Kategorie)
	}
	for _, v := range basiswerte.AlleKampftechniken {
		h.NewKampftechnik(v.Name, v.Typ == "Fernkampf", v.Typ == "NahkampfNurAT", v.Leiteigenschaft, v.Steigerungsfaktor)
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
			tmpVorteil := basiswerte.VorUndNachteil{Name: vorteil.Name, Vorraussetzungen: vorteil.Vorraussetzungen, APKosten: vorteil.APKosten, FromSpezies: true}
			h.Vorteile = append(h.Vorteile, tmpVorteil)
			//h.Vorteile = append(h.Vorteile, vorteil)
		}
	}

	for _, v := range h.Spezies.Nachteile {
		nachteil := basiswerte.GetNachteil(v)
		if nachteil != nil {
			tmpNachteil := basiswerte.VorUndNachteil{Name: nachteil.Name, Vorraussetzungen: nachteil.Vorraussetzungen, APKosten: nachteil.APKosten, FromSpezies: true}
			h.Nachteile = append(h.Nachteile, tmpNachteil)
			//h.Nachteile = append(h.Nachteile, nachteil)
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

func (h *Held) SetProfession(profession *basiswerte.Profession) error {
	h.Profession = *profession
	for _, v := range h.Profession.Talente {
		fmt.Println(v)
		talent := v[0]
		t_num, err := strconv.Atoi(v[1])
		if err != nil {
			return errors.New("Error converting number in profession! Check regeln/profession for bugs in data.")
		}
		h.Talente.Get(talent).AddValue(t_num)
	}
	for _, v := range h.Profession.Kampftechniken {
		kt := v[0]
		k_num, err := strconv.Atoi(v[1])
		if err != nil {
			return errors.New("Error converting number in profession! Check regeln/profession for bugs in data.")
		}
		h.Kampftechniken.Get(kt).AddValue(k_num - 6)
	}
	for _, v := range h.Profession.Zauber {
		fmt.Println(v)
		kt := v[0]
		k_num, err := strconv.Atoi(v[1])
		if err != nil {
			return errors.New("Error converting number in profession! Check regeln/profession for bugs in data.")
		}
		_, exists := basiswerte.AlleZauber[kt]
		if !exists {
			return errors.New("Unknown Zauber! Check regeln/profession for bugs in data.")
		}
		zauber, _ := basiswerte.AlleZauber[kt]
		h.NewZauber(&zauber)
		z := h.Zauber.Get(kt)
		if z != nil {
			z.SetMaxErschaffung(h.Erfahrungsgrad.Fertigkeit)
			z.AddValue(k_num)
		}
	}
	for _, v := range h.Profession.Liturgien {
		fmt.Println(v)
		kt := v[0]
		k_num, err := strconv.Atoi(v[1])
		if err != nil {
			return errors.New("Error converting number in profession! Check regeln/profession for bugs in data.")
		}
		_, exists := basiswerte.AlleLiturgien[kt]
		if !exists {
			return errors.New("Unknown Liturgie! Check regeln/profession for bugs in data.")
		}
		liturgie, _ := basiswerte.AlleLiturgien[kt]
		h.NewLiturgie(&liturgie)
		l := h.Liturgien.Get(kt)
		if l != nil {
			l.SetMaxErschaffung(h.Erfahrungsgrad.Fertigkeit)
			l.AddValue(k_num)
		}
	}

	for _, v := range h.Profession.Sonderfertigkeiten {
		var bereich *[]*basiswerte.Sonderfertigkeit
		switch basiswerte.GetSFType(v) {
		case basiswerte.ALLGEMEIN:
			{
				bereich = &h.Sonderfertigkeiten.Allgemeine
			}
		case basiswerte.KARMAL:
			{
				bereich = &h.Sonderfertigkeiten.Karmale
			}
		case basiswerte.MAGISCH:
			{
				bereich = &h.Sonderfertigkeiten.Magische
			}
		case basiswerte.KAMPF:
			{
				bereich = &h.Sonderfertigkeiten.Kampf
			}
		case basiswerte.SPRACHE:
			{
				bereich = &h.Sonderfertigkeiten.Sprachen
			}
		case basiswerte.SCHRIFT:
			{
				bereich = &h.Sonderfertigkeiten.Schriften
			}
		case basiswerte.UNBEKANNT:
			{
				bereich = nil
			}
		}
		if bereich == nil {
			fmt.Printf("Unbekannte Sonderfertigkeit %s! Überspringe!\n", v)
		}
		sf := basiswerte.GetSF(v)
		if sf != nil {
			found := false
			for _, v := range *bereich {
				if v.Name == sf.Name {
					found = true
				}
			}
			if !found {
				*bereich = append(*bereich, sf)
			}
		}
	}
	h.APAusgeben(h.Profession.APKosten)
	return nil
}

func (h *Held) APGesamt() int { return (h.AP + h.AP_spent) }

func (h *Held) APAusgeben(menge int) {
	h.AP -= menge
	h.AP_spent += menge
}

func (h *Held) IsMagisch() bool {
	for _, v := range h.Vorteile {
		if v.Name == "Zauberer" {
			return true
		}
	}
	return false
}

func (h *Held) IsKarmal() bool {
	for _, v := range h.Vorteile {
		if v.Name == "Geweihter" {
			return true
		}
	}
	return false
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

func (h *Held) NewKampftechnik(name string, isFernkampf bool, isNurAT bool, leiteigenschaften []string, sf string) bool {
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
		h.Kampftechniken.Add(basiswerte.MakeNahkampf(name, 6, lt, h.Eigenschaften.Eigenschaften["MU"], sf, isNurAT))
	}
	return h.Kampftechniken.Exists(name)
}

func (h *Held) NewZauber(zauber *basiswerte.ZauberType) bool {
	if h.Zauber.Exists(zauber.Name) {
		return false
	}
	var e1 *basiswerte.Eigenschaft
	if zauber.Probe.Eigenschaften[0] != "-" {
		e1 = h.Eigenschaften.Eigenschaften[zauber.Probe.Eigenschaften[0]]
	} else {
		e1 = nil
	}

	var e2 *basiswerte.Eigenschaft
	if zauber.Probe.Eigenschaften[1] != "-" {
		e2 = h.Eigenschaften.Eigenschaften[zauber.Probe.Eigenschaften[1]]
	} else {
		e1 = nil
	}

	var e3 *basiswerte.Eigenschaft
	if zauber.Probe.Eigenschaften[2] != "-" {
		e3 = h.Eigenschaften.Eigenschaften[zauber.Probe.Eigenschaften[2]]
	} else {
		e1 = nil
	}
	fmt.Println(e1, e2, e3)
	h.Zauber.Add(basiswerte.MakeZauber(zauber.Name, 0, e1, e2, e3, zauber))

	return h.Zauber.Exists(zauber.Name)
}

func (h *Held) NewLiturgie(liturgie *basiswerte.LiturgieType) bool {
	if h.Liturgien.Exists(liturgie.Name) {
		return false
	}
	var e1 *basiswerte.Eigenschaft
	if liturgie.Probe.Eigenschaften[0] != "-" {
		e1 = h.Eigenschaften.Eigenschaften[liturgie.Probe.Eigenschaften[0]]
	} else {
		e1 = nil
	}

	var e2 *basiswerte.Eigenschaft
	if liturgie.Probe.Eigenschaften[1] != "-" {
		e2 = h.Eigenschaften.Eigenschaften[liturgie.Probe.Eigenschaften[1]]
	} else {
		e1 = nil
	}

	var e3 *basiswerte.Eigenschaft
	if liturgie.Probe.Eigenschaften[2] != "-" {
		e3 = h.Eigenschaften.Eigenschaften[liturgie.Probe.Eigenschaften[2]]
	} else {
		e1 = nil
	}
	fmt.Println(e1, e2, e3)
	h.Liturgien.Add(basiswerte.MakeLiturgie(liturgie.Name, 0, e1, e2, e3, liturgie))

	return h.Liturgien.Exists(liturgie.Name)
}
