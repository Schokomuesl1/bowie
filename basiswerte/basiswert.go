package basiswerte

import (
    "fmt"
)

type BerechneterWert struct {
    Name          string
    BaseValues    []ValueType
    Mult          float32
    dependingOnMe dependenceStorage
}

func MakeBerechneterWert(name string, mult float32, base []DependentValue) *BerechneterWert {
    var cal BerechneterWert
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

func (c *BerechneterWert) Value() int {
    var sum int
    for _, v := range c.BaseValues {
        sum += v.Value()
    }
    return int((float32(sum) * c.Mult) + 0.5)
}

func (c BerechneterWert) Update() {
    c.dependingOnMe.NotifyValueChanged() // no re-calculation, this happens in Value(). Only notify we might have to be recalculated
}

func (c BerechneterWert) String() string {
    return fmt.Sprintf("%s: %d", c.Name, c.Value())
}

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
        Astralenergie:   MakeBasiswert("Astralenergie", 0, []DependentValue{}, 0),
        Karmaenergie:    MakeBasiswert("Karmaenergie", 0, []DependentValue{}, 0),
        Seelenkraft:     MakeBasiswert("Seelenkraft", (1.0 / 6.0), []DependentValue{eigenschaften.Get("MU"), eigenschaften.Get("KL"), eigenschaften.Get("IN")}, spezies.SK),
        Zaehigkeit:      MakeBasiswert("Zaehigkeit", (1.0 / 6.0), []DependentValue{eigenschaften.Get("KO"), eigenschaften.Get("KO"), eigenschaften.Get("KK")}, spezies.ZK),
        Ausweichen:      MakeBasiswert("Ausweichen", 0.5, []DependentValue{eigenschaften.Get("GE")}, 0),
        Initiative:      MakeBasiswert("Initiative", 0.5, []DependentValue{eigenschaften.Get("GE"), eigenschaften.Get("MU")}, 0),
        Geschwindigkeit: MakeBasiswert("Geschwindigkeit", 1, []DependentValue{}, spezies.GS)}
}
