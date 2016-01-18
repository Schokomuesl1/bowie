package bowie

import (
	"fmt"
)

type DependentValue interface {
	Register(DependingValue)
	NotifyValueChanged()
	Value() int
}

type DependingValue interface {
	Update()
}

type ValueType interface {
	Value() int
}

type ErgebnisTyp int

const (
	ERFOLG ErgebnisTyp = iota
	MISSERFOLG
	PATZER
	GLUECKLICH
	DOPPEL_EINS
	DREIFACH_EINS
	DOPPEL_ZWANZIG
	DREIFACH_ZWANZIG
)

func (e ErgebnisTyp) String() string {
	switch e {
	case ERFOLG:
		return "Erfolg"
	case MISSERFOLG:
		return "Misserfolg"
	case PATZER:
		return "Patzer"
	case GLUECKLICH:
		return "Gluecklich"
	case DOPPEL_EINS:
		return "Doppel-Eins"
	case DREIFACH_EINS:
		return "Dreifach-Eins"
	case DOPPEL_ZWANZIG:
		return "Doppel-Zwanzig"
	case DREIFACH_ZWANZIG:
		return "Dreifach-Zwanzig"
	default:
		return "Fehler!"
	}
}

type WurfUndZiel struct {
	Wurf int
	Ziel int
}

func (w WurfUndZiel) String() string { return fmt.Sprintf("%d(%d)", w.Wurf, w.Ziel) }

type ProbenErgebnis struct {
	Ergebnis  []WurfUndZiel
	Mod       int
	Differenz int
	Typ       ErgebnisTyp
}

func MakeEinzelErgebnis(wurf int, ziel int, mod int, diff int, typ ErgebnisTyp) *ProbenErgebnis {
	var pe ProbenErgebnis
	pe.Ergebnis = make([]WurfUndZiel, 1)
	pe.Ergebnis[0] = WurfUndZiel{wurf, ziel}
	pe.Differenz = diff
	pe.Typ = typ
	return &pe
}

func MakeMehrfachErgebnis(ergebnis []WurfUndZiel, mod int, diff int, typ ErgebnisTyp) *ProbenErgebnis {
	var pe ProbenErgebnis
	pe.Ergebnis = make([]WurfUndZiel, len(ergebnis))
	for i, v := range ergebnis {
		pe.Ergebnis[i] = v
	}
	pe.Differenz = diff
	pe.Typ = typ
	return &pe
}

func (p ProbenErgebnis) String() string {
	return fmt.Sprintf("Roll(s): %s, Mod: %d, Diff: %d, Ergebnis: %s", p.Ergebnis, p.Mod, p.Differenz, p.Typ)
}

type ProbenType interface {
	Probe() (bool, ProbenErgebnis)
	ProbeMod(int) (bool, ProbenErgebnis)
}

type dependenceStorage struct {
	dependsOnMe []DependingValue
}

func (s dependenceStorage) Register(d DependingValue) {
	// first check if d is already registered
	for _, v := range s.dependsOnMe {
		if d == v {
			return
		}
	}
	n := len(s.dependsOnMe)
	if n == cap(s.dependsOnMe) {
		newDependsOnMe := make([]DependingValue, len(s.dependsOnMe), 2*len(s.dependsOnMe)+1)
		copy(newDependsOnMe, s.dependsOnMe)
		s.dependsOnMe = newDependsOnMe
	}
	s.dependsOnMe = s.dependsOnMe[0 : n+1]
	s.dependsOnMe[n] = d
}

func (s dependenceStorage) NotifyValueChanged() {
	for _, v := range s.dependsOnMe {
		v.Update()
	}
}

type Eigenschaft struct {
	Name          string
	Wert          int
	dependingOnMe dependenceStorage
}

func MakeEigenschaft(name string, wert int) *Eigenschaft {
	var e Eigenschaft
	e.Name = name
	e.Wert = wert
	return &e
}

func (e *Eigenschaft) Set(value int) {
	e.Wert = value
}

func (e *Eigenschaft) Register(d DependingValue) {
	e.dependingOnMe.Register(d)
}

func (e *Eigenschaft) NotifyValueChanged() {
	e.dependingOnMe.NotifyValueChanged()
}

func (e *Eigenschaft) Increment() {
	e.Wert++
	e.dependingOnMe.NotifyValueChanged()
}

func (e *Eigenschaft) Value() int {
	return e.Wert
}

func (e *Eigenschaft) Probe() (bool, ProbenErgebnis) {
	return e.ProbeMod(0)
}

//func MakeEinzelErgebnis(wurf int, ziel int, diff int, typ ErgebnisTyp) *ProbenErgebnis {
func (e *Eigenschaft) ProbeMod(mod int) (bool, ProbenErgebnis) {
	roll := int(NormalDice{Sides: 20}.Roll())
	diff := e.Wert - mod - roll
	switch roll {
	case 20:
		// try again...
		roll2 := int(NormalDice{Sides: 20}.Roll())
		diff2 := e.Wert - mod - roll
		if roll2 == 20 || diff2 < 0 {
			return false, *MakeEinzelErgebnis(roll, e.Wert, mod, diff, PATZER)
		} else {
			return false, *MakeEinzelErgebnis(roll, e.Wert, mod, diff, MISSERFOLG)
		}
	case 1:
		// try again...
		roll2 := int(NormalDice{Sides: 20}.Roll())
		diff2 := e.Wert - mod - roll
		if roll2 == 1 || diff2 >= 0 {
			return true, *MakeEinzelErgebnis(roll, e.Wert, mod, diff, GLUECKLICH)
		} else {
			return true, *MakeEinzelErgebnis(roll, e.Wert, mod, diff, ERFOLG)
		}
	default:
		if diff >= 0 {
			return true, *MakeEinzelErgebnis(roll, e.Wert, mod, diff, ERFOLG)
		} else {
			return false, *MakeEinzelErgebnis(roll, e.Wert, mod, diff, MISSERFOLG)
		}
	}
}

func (e *Eigenschaft) String() string {
	return fmt.Sprintf("%s(%d)", e.Name, e.Wert)
}

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

// Talent CTor
func MakeTalent(name string, wert int, e1 *Eigenschaft, e2 *Eigenschaft, e3 *Eigenschaft) *Talent {
	//tal := Talent{name, wert, {e1, e2, e3}, dependenceStorage{}}
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
	switch anz1 {
	case 2:
		return true, *MakeMehrfachErgebnis(gesamtErgebnis[:], mod, gesamtVerbraucht, DOPPEL_EINS)
	case 3:
		return true, *MakeMehrfachErgebnis(gesamtErgebnis[:], mod, gesamtVerbraucht, DREIFACH_EINS)
	}
	switch anz20 {
	case 2:
		return false, *MakeMehrfachErgebnis(gesamtErgebnis[:], mod, gesamtVerbraucht, DOPPEL_ZWANZIG)
	case 3:
		return false, *MakeMehrfachErgebnis(gesamtErgebnis[:], mod, gesamtVerbraucht, DREIFACH_ZWANZIG)
	}
	if gesamtVerbraucht < t.Wert-mod {
		return false, *MakeMehrfachErgebnis(gesamtErgebnis[:], mod, gesamtVerbraucht, MISSERFOLG)
	} else {
		return true, *MakeMehrfachErgebnis(gesamtErgebnis[:], mod, gesamtVerbraucht, ERFOLG)
	}
}

func (t *Talent) String() string {
	return fmt.Sprintf("%s: %d - (%s/%s/%s)", t.Name, t.Value(), t.Eigenschaften[0], t.Eigenschaften[1], t.Eigenschaften[2])
}

type CalculatedDependentValue struct {
	Name          string
	BaseValues    []ValueType
	Mult          float32
	dependingOnMe dependenceStorage
}

func MakeCalculatedDependentValue(name string, mult float32, base []DependentValue) *CalculatedDependentValue {
	var cal CalculatedDependentValue
	cal.Name = name
	cal.Mult = mult
	if cal.Mult == 0 {
		cal.Mult = 1
	}
	cal.BaseValues = make([]ValueType, len(base))
	for i, v := range base {
		cal.BaseValues[i] = v
		v.Register(cal)
	}
	return &cal
}

func (c *CalculatedDependentValue) Value() int {
	var sum int
	for _, v := range c.BaseValues {
		sum += v.Value()
	}
	return int(float32(sum) * c.Mult)
}

func (c CalculatedDependentValue) Update() {
	c.dependingOnMe.NotifyValueChanged() // no re-calculation, this happens in Value(). Only notify we might have to be recalculated
}

func (c CalculatedDependentValue) String() string {
	return fmt.Sprintf("%s: %d", c.Name, c.Value())
}
