package basiswerte

import (
	"fmt"
)

type Eigenschaft struct {
	Name          string
	Wert          int
	max           int
	min           int
	dependingOnMe dependenceStorage
}

func MakeEigenschaft(name string, wert int) *Eigenschaft {
	var e Eigenschaft
	e.Name = name
	e.Wert = wert
	return &e
}

func (e *Eigenschaft) Add(value int) bool {
	return e.Set(e.Wert + value)
}

func (e *Eigenschaft) Set(value int) bool {
	if value <= e.max && value >= e.min {
		e.Wert = value
		return true
	}
	return false
}

func (e *Eigenschaft) SetWithRange(value int, min int, max int) bool {
	fmt.Println("SetWithRange", value, min, max)
	if value <= max && value >= min {
		e.Wert = value
		e.SetMin(min)
		e.SetMax(max)
		fmt.Println("done setting ", e)
		return true
	}
	fmt.Println("not setting ", e)
	return false
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

func (e *Eigenschaft) Value() int                    { return e.Wert }
func (e *Eigenschaft) Max() int                      { return e.max }
func (e *Eigenschaft) Min() int                      { return e.min }
func (e *Eigenschaft) SetMin(m int)                  { e.min = m }
func (e *Eigenschaft) SetMax(m int)                  { e.max = m }
func (e *Eigenschaft) Probe() (bool, ProbenErgebnis) { return e.ProbeMod(0) }
func (e *Eigenschaft) KannSteigern() string {
	if e.Value()+1 <= e.Max() {
		return ""
	}
	return "disabled"
}

func (e *Eigenschaft) KannSenken() string {
	if e.Value()-1 >= e.Min() {
		return ""
	}
	return "disabled"
}

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
	return fmt.Sprintf("%2s(%2d) [%2d,%2d]", e.Name, e.Wert, e.min, e.max)
}

type EigenschaftHandler struct {
	Eigenschaften map[string]*Eigenschaft
}

func (e *EigenschaftHandler) Init(name string, eigenschaftMax int) bool {
	if !e.Exists(name) {
		e.Add(name)
	}
	e.Eigenschaften[name].SetWithRange(8, 8, eigenschaftMax)
	return true
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
