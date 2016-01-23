package basiswerte

import (
	"fmt"
)

type Talent struct {
	Name          string
	Wert          int
	Eigenschaften [3]*Eigenschaft
	dependingOnMe dependenceStorage
}

func (t Talent) Register(d DependingValue) {
	t.dependingOnMe.Register(d)
}

func (t Talent) NotifyValueChanged() {
	t.dependingOnMe.NotifyValueChanged()
}

func (t *Talent) Update() {
	// one of the Eigenschaften changed :-)
	for _, e := range t.Eigenschaften {
		fmt.Println(e.Name, e.Wert)
	}
}

func (t *Talent) AddValue(value int) {
	t.Wert += value
}

func (t *Talent) Increment() {
	t.Wert++
}

func (t *Talent) Decrement() {
	t.Wert++
}

// Talent CTor
func MakeTalent(name string, wert int, e1 *Eigenschaft, e2 *Eigenschaft, e3 *Eigenschaft) *Talent {
	var tal Talent
	tal.Name = name
	tal.Wert = wert
	tal.Eigenschaften[0] = e1
	tal.Eigenschaften[1] = e2
	tal.Eigenschaften[2] = e3
	for _, e := range tal.Eigenschaften {
		e.dependingOnMe.Register(&tal)
	}
	return &tal
}

func (t *Talent) Value() int {
	return t.Wert
}

func (t *Talent) Probe() (bool, ProbenErgebnis) {
	return t.ProbeMod(0)
}

func (t *Talent) ProbeMod(mod int) (bool, ProbenErgebnis) {
	var result [3]ProbenErgebnis
	for i, v := range t.Eigenschaften {
		_, result[i] = v.ProbeMod(mod)
	}
	anz1 := 0
	anz20 := 0
	gesamtVerbraucht := 0
	var gesamtErgebnis [3]WurfUndZiel
	for i, v := range result {
		gesamtErgebnis[i] = v.Ergebnis[0]
		switch v.Ergebnis[0].Wurf {
		case 1:
			anz1++
		case 20:
			anz20++
		}
		if v.Differenz > 0 {
			gesamtVerbraucht += v.Differenz
		}
	}

	success := false
	ergebnis := ERFOLG
	switch anz1 {
	case 2:
		success = true
		ergebnis = DOPPEL_EINS
	case 3:
		success = true
		ergebnis = DREIFACH_EINS
	}
	switch anz20 {
	case 2:
		success = false
		ergebnis = DOPPEL_ZWANZIG
	case 3:
		success = false
		ergebnis = DREIFACH_ZWANZIG
	}
	if gesamtVerbraucht < t.Wert-mod {
		success = false
		ergebnis = MISSERFOLG
	} else {
		success = true
		ergebnis = ERFOLG
	}
	return success, *MakeMehrfachErgebnis(gesamtErgebnis[:], mod, gesamtVerbraucht, ergebnis)
}

func (t *Talent) String() string {
	return fmt.Sprintf("%25s: %2d - (%s/%s/%s)", t.Name, t.Value(), t.Eigenschaften[0], t.Eigenschaften[1], t.Eigenschaften[2])
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
