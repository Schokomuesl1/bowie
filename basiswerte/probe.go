package basiswerte

import (
    "fmt"
)

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

func (w WurfUndZiel) String() string { return fmt.Sprintf("%2d(%2d)", w.Wurf, w.Ziel) }

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
