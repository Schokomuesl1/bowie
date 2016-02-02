package basiswerte

import (
	//	"fmt"
	"sort"
)

type Liturgie struct {
	Name           string
	Wert           int
	maxErschaffung int
	Eigenschaften  [3]*Eigenschaft
	dependingOnMe  dependenceStorage
	Weiteres       *LiturgieType
}

func (l Liturgie) Register(d DependingValue) {
	l.dependingOnMe.Register(d)
}

func (l Liturgie) NotifyValueChanged() {
	l.dependingOnMe.NotifyValueChanged()
}

func (l *Liturgie) Update() {
	// one of the Eigenschaften changed :-)
}

func (l *Liturgie) AddValue(value int) {
	tmp := l.Wert + value
	if tmp >= l.Min() && tmp <= l.Max() {
		l.Wert += value
	}
}

func (l *Liturgie) SetMaxErschaffung(max int) { l.maxErschaffung = max }
func (l *Liturgie) Increment() {
	l.AddValue(1)
}

func (l *Liturgie) Decrement() {
	l.AddValue(-1)
}

func (l Liturgie) SK() string { return l.Weiteres.Steigerungsfaktor }

func (l *Liturgie) Value() int { return l.Wert }
func (l *Liturgie) Min() int   { return 0 }
func (l *Liturgie) Max() int {
	max := 0
	for _, v := range l.Eigenschaften {
		if v == nil {
			continue
		}
		if v.Value() > max {
			max = v.Value()
		}
	}
	max += 2 //  höchste eigenschaft +2
	if max < l.maxErschaffung {
		return max
	}
	return l.maxErschaffung
}

func (l *Liturgie) KannSteigern() string {
	if l.Weiteres.Steigerungsfaktor == "-" {
		// special case: segnungen
		return "disabled"
	}
	if l.Value()+1 <= l.Max() {
		return ""
	}
	return "disabled"
}

func (l *Liturgie) KannSenken() string {
	if l.Weiteres.Steigerungsfaktor == "-" {
		// special case: segnungen
		return "disabled"
	}
	if l.Value()-1 >= l.Min() {
		return ""
	}
	return "disabled"
}

func MakeLiturgie(name string, wert int, e1 *Eigenschaft, e2 *Eigenschaft, e3 *Eigenschaft, ref *LiturgieType) *Liturgie {
	var lit Liturgie
	lit.Name = name
	lit.Wert = wert
	lit.Eigenschaften[0] = e1
	lit.Eigenschaften[1] = e2
	lit.Eigenschaften[2] = e3
	lit.Weiteres = ref
	for _, e := range lit.Eigenschaften {
		e.dependingOnMe.Register(&lit)
	}
	return &lit
}

type LiturgieHandler struct {
	Liturgien map[string]*Liturgie
}

func NewLiturgieHandler() *LiturgieHandler {
	return &LiturgieHandler{Liturgien: make(map[string]*Liturgie)}
}

func (l *LiturgieHandler) Exists(name string) bool {
	_, existing := l.Liturgien[name]
	return existing
}

func (l *LiturgieHandler) Add(liturgie *Liturgie) bool {
	if l.Exists(liturgie.Name) {
		return false
	}
	l.Liturgien[liturgie.Name] = liturgie
	return true
}

func (l *LiturgieHandler) Get(liturgie string) *Liturgie {
	if !l.Exists(liturgie) {
		return nil
	}
	return l.Liturgien[liturgie]
}

func (l *LiturgieHandler) SetErschaffungsMax(max int) {
	for _, v := range l.Liturgien {
		v.SetMaxErschaffung(max)
	}
}

type LiturgieListe []*Liturgie

func (l LiturgieListe) Len() int           { return len(l) }
func (l LiturgieListe) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l LiturgieListe) Less(i, j int) bool { return l[i].Name < l[j].Name }

func (l *LiturgieHandler) Sortiert() (ll LiturgieListe) {
	ll = make(LiturgieListe, 0)
	for _, v := range l.Liturgien {
		ll = append(ll, v)
	}
	sort.Sort(ll)
	return ll
}
