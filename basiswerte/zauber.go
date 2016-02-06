package basiswerte

import (
	"fmt"
	"sort"
)

type Zauber struct {
	Name          string
	Wert          int
	Eigenschaften [3]*Eigenschaft
	dependingOnMe dependenceStorage
	Weiteres      *ZauberType
}

func (z Zauber) Register(d DependingValue) {
	z.dependingOnMe.Register(d)
}

func (z Zauber) NotifyValueChanged() {
	z.dependingOnMe.NotifyValueChanged()
}

func (z *Zauber) Update() {
	// one of the Eigenschaften changed :-)
}

func (z *Zauber) AddValue(value int) {
	tmp := z.Wert + value
	if tmp >= z.Min() && tmp <= z.Max() {
		z.Wert += value
	}
}

func (z *Zauber) Increment() {
	z.AddValue(1)
}

func (z *Zauber) Decrement() {
	z.AddValue(-1)
}

func (z *Zauber) SK() string { return z.Weiteres.Steigerungsfaktor }

func (z *Zauber) Value() int { return z.Wert }
func (z *Zauber) Min() int   { return 0 }
func (z *Zauber) Max() int {
	max := 0
	for _, v := range z.Eigenschaften {
		fmt.Println(v)
		if v == nil {
			continue
		}
		if v.Value() > max {
			max = v.Value()
		}
	}
	return max + 2
}

func (z *Zauber) KannSteigern() string {
	if z.Weiteres.Steigerungsfaktor == "-" {
		// special case: zaubertricks
		return "disabled"
	}
	fmt.Println(z.Value()+1, z.Max())
	if z.Value()+1 <= z.Max() {
		return ""
	}
	return "disabled"
}

func (z *Zauber) KannSenken() string {
	if z.Weiteres.Steigerungsfaktor == "-" {
		// special case: zaubertricks
		return "disabled"
	}
	if z.Value()-1 >= z.Min() {
		return ""
	}
	return "disabled"
}

func MakeZauber(name string, wert int, e1 *Eigenschaft, e2 *Eigenschaft, e3 *Eigenschaft, ref *ZauberType) *Zauber {
	var zau Zauber
	zau.Name = name
	zau.Wert = wert
	zau.Eigenschaften[0] = e1
	zau.Eigenschaften[1] = e2
	zau.Eigenschaften[2] = e3
	zau.Weiteres = ref
	for _, e := range zau.Eigenschaften {
		e.dependingOnMe.Register(&zau)
	}
	return &zau
}

type ZauberHandler struct {
	Zaubers map[string]*Zauber
}

func NewZauberHandler() *ZauberHandler {
	return &ZauberHandler{Zaubers: make(map[string]*Zauber)}
}

func (z *ZauberHandler) Exists(name string) bool {
	_, existing := z.Zaubers[name]
	return existing
}

func (z *ZauberHandler) Add(Zauber *Zauber) bool {
	if z.Exists(Zauber.Name) {
		return false
	}
	z.Zaubers[Zauber.Name] = Zauber
	return true
}

func (z *ZauberHandler) Get(Zauber string) *Zauber {
	if !z.Exists(Zauber) {
		return nil
	}
	return z.Zaubers[Zauber]
}

type ZauberListe []*Zauber

func (z ZauberListe) Len() int           { return len(z) }
func (z ZauberListe) Swap(i, j int)      { z[i], z[j] = z[j], z[i] }
func (z ZauberListe) Less(i, j int) bool { return z[i].Name < z[j].Name }

func (z *ZauberHandler) Sortiert() (zl ZauberListe) {
	zl = make(ZauberListe, 0)
	for _, v := range z.Zaubers {
		zl = append(zl, v)
	}
	sort.Sort(zl)
	return zl
}
