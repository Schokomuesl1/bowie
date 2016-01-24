package basiswerte

import (
	"fmt"
)

type KampftechnikType struct {
	Name              string
	Typ               string
	Leiteigenschaft   []string
	Steigerungsfaktor string
}

//"Name": "Armbrüste"
//"Typ": "Fernkampf"
//"Leiteigenschaft": ["FF"]
//"Steigerungsfaktor":"B"}

type Kampftechnik struct {
	Name              string
	Wert              int
	IsFernkampf       bool
	Leiteigenschaften []*Eigenschaft
	SK                string
	mut               *Eigenschaft
	ff                *Eigenschaft
	dependingOnMe     dependenceStorage
}

func (k Kampftechnik) Register(d DependingValue) {
	k.dependingOnMe.Register(d)
}

func (k *Kampftechnik) Update() {
	// nothing here
}

func (k Kampftechnik) NotifyValueChanged() {
	k.dependingOnMe.NotifyValueChanged()
}

func (k *Kampftechnik) AddValue(value int) {
	tmp := k.Wert + value
	if tmp >= k.Min() && tmp <= k.Max() {
		k.Wert += value
	}
}

func (k *Kampftechnik) AT() int {
	if k.IsFernkampf {
		return 0
	}
	extra := k.mut.Value() - 8
	if extra > 0 {
		return k.Wert + (extra / 3)
	}
	return k.Wert
}

func (k *Kampftechnik) PA() int {
	if k.IsFernkampf {
		return 0
	}
	extra := 0
	for _, v := range k.Leiteigenschaften {
		if v.Value() > extra {
			extra = v.Value()
		}
	}
	if extra > 0 {
		return int((float32(k.Wert)/2.0)+0.5) + (extra / 3)
	}
	return int((float32(k.Wert) / 2.0) + 0.5)
}

func (k *Kampftechnik) FK() int {
	if !k.IsFernkampf {
		return 0
	}
	extra := k.ff.Value() - 8
	if extra > 0 {
		return k.Wert + (extra / 3)
	}
	return k.Wert
}

func (k *Kampftechnik) Value() int { return k.Wert }
func (k *Kampftechnik) Min() int   { return 6 } // by default
func (k *Kampftechnik) Max() int {
	max := 0
	for _, v := range k.Leiteigenschaften {
		if v.Value() > max {
			max = v.Value()
		}
	}
	return max + 2 //  höchste eigenschaft +2
}

func (k *Kampftechnik) KannSteigern() string {
	if k.Value()+1 <= k.Max() {
		return ""
	}
	return "disabled"
}

func (k *Kampftechnik) KannSenken() string {
	if k.Value()-1 >= k.Min() {
		return ""
	}
	return "disabled"
}

func (k *Kampftechnik) String() string {
	if k.IsFernkampf {
		return fmt.Sprintf("%25s: %2d [%2d,%2d] - FK: %2d", k.Name, k.Value(), k.Min(), k.Max(), k.FK())
	}
	return fmt.Sprintf("%25s: %2d [%2d,%2d] - AT/PA: %2d/%2d", k.Name, k.Value(), k.Min(), k.Max(), k.AT(), k.PA())
}

func MakeNahkampf(name string, wert int, leiteigenschaften []*Eigenschaft, mut *Eigenschaft, sf string) *Kampftechnik {
	var kt Kampftechnik
	kt.Name = name
	kt.Wert = wert
	kt.IsFernkampf = false
	kt.Leiteigenschaften = make([]*Eigenschaft, len(leiteigenschaften))
	//	kt.Leiteigenschaften = leiteigenschaften
	for i, v := range leiteigenschaften {
		kt.Leiteigenschaften[i] = v
	}
	kt.SK = sf
	for _, e := range kt.Leiteigenschaften {
		e.dependingOnMe.Register(&kt)
	}
	kt.ff = nil
	kt.mut = mut
	mut.dependingOnMe.Register(&kt)
	return &kt
}

func MakeFernkampf(name string, wert int, leiteigenschaften []*Eigenschaft, fingerfertigkeit *Eigenschaft, sf string) *Kampftechnik {
	var kt Kampftechnik
	fmt.Println("kt.Name = name")
	kt.Name = name
	fmt.Println("kt.Wert = wert")
	kt.Wert = wert
	fmt.Println("kt.IsFernkampf = true")
	kt.IsFernkampf = true
	fmt.Println("kt.Leiteigenschaften = make([]*Eigenschaft, len(leiteigenschaften))")
	kt.Leiteigenschaften = make([]*Eigenschaft, len(leiteigenschaften))
	//kt.Leiteigenschaften = leiteigenschaften
	for i, v := range leiteigenschaften {
		kt.Leiteigenschaften[i] = v
	}
	fmt.Println("kt.SK = sf")
	kt.SK = sf
	for _, e := range kt.Leiteigenschaften {
		e.dependingOnMe.Register(&kt)
	}
	kt.ff = fingerfertigkeit
	kt.mut = nil
	fingerfertigkeit.dependingOnMe.Register(&kt)
	return &kt
}

type KampftechnikHandler struct {
	Kampftechniken map[string]*Kampftechnik
}

func NewKampftechnikHandler() *KampftechnikHandler {
	return &KampftechnikHandler{Kampftechniken: make(map[string]*Kampftechnik)}
}

func (k *KampftechnikHandler) Exists(name string) bool {
	_, existing := k.Kampftechniken[name]
	return existing
}

func (k *KampftechnikHandler) Add(Kampftechnik *Kampftechnik) bool {
	if k.Exists(Kampftechnik.Name) {
		return false
	}
	k.Kampftechniken[Kampftechnik.Name] = Kampftechnik
	return true
}

func (k *KampftechnikHandler) Get(name string) *Kampftechnik {
	if !k.Exists(name) {
		return nil
	}
	return k.Kampftechniken[name]
}
