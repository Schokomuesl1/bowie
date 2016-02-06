package erschaffung

import (
	//	"fmt"
	"github.com/Schokomuesl1/bowie/basiswerte"
	"github.com/Schokomuesl1/bowie/held"
	"strconv"
)

func ErschaffeHeld(erfahrungsgrad string) (*held.Held, *ErschaffungsValidator) {
	h := held.NewHeld()
	e := MakeErschaffungsValidator(h, erfahrungsgrad)
	if e == nil {
		return nil, nil
	}
	h.AP = e.Grad.AP
	return h, e
}

func SFAvailable(Held *held.Held, SF *basiswerte.Sonderfertigkeit) (result bool) {
	result = false
	// talente checken
	for _, v := range SF.Vorraussetzungen.Talente {
		t := Held.Talente.Get(v[0])
		v_num, err := strconv.Atoi(v[1])
		if t == nil || err != nil || t.Value() < v_num {
			return
		}
	}

	// Eigenschaften checken
	for _, v := range SF.Vorraussetzungen.Eigenschaften {
		t := Held.Eigenschaften.Get(v[0])
		v_num, err := strconv.Atoi(v[1])
		if t == nil || err != nil || t.Value() < v_num {
			return
		}
	}

	// Vorteile...
	for _, v := range SF.Vorraussetzungen.Vorteile {
		found := false
		for _, k := range Held.Vorteile {
			if k.Name == v {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}
	// Verbotene Nachteile...
	for _, v := range SF.Vorraussetzungen.NichtNachteil {
		for _, k := range Held.Vorteile {
			if k.Name == v {
				return
			}
		}
	}

	for _, v := range SF.Vorraussetzungen.Sonderfertigkeiten {
		found := false
		for _, k := range Held.Sonderfertigkeiten.Allgemeine {
			if k.Name == v {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}
	for _, v := range SF.Vorraussetzungen.Sonderfertigkeiten {
		found := false
		for _, k := range Held.Sonderfertigkeiten.Karmale {
			if k.Name == v {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}
	for _, v := range SF.Vorraussetzungen.Sonderfertigkeiten {
		found := false
		for _, k := range Held.Sonderfertigkeiten.Magische {
			if k.Name == v {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}
	for _, v := range SF.Vorraussetzungen.Sonderfertigkeiten {
		found := false
		for _, k := range Held.Sonderfertigkeiten.Kampf {
			if k.Name == v {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}
	return true
}

func VorUndNachteilAvailable(Held *held.Held, VorOderNachteil *basiswerte.VorUndNachteil) (result bool) {
	result = false
	//fmt.Println(VorOderNachteil)
	for _, v := range VorOderNachteil.Vorraussetzungen.Vorteile {
		//fmt.Println(v)
		found := false
		for _, k := range Held.Vorteile {
			if v == k.Name {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}

	for _, v := range VorOderNachteil.Vorraussetzungen.Nachteile {
		found := false
		for _, k := range Held.Nachteile {
			if v == k.Name {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}

	for _, v := range VorOderNachteil.Vorraussetzungen.NichtVorteile {
		for _, k := range Held.Vorteile {
			if v == k.Name {
				return
			}
		}
	}

	for _, v := range VorOderNachteil.Vorraussetzungen.NichtNachteile {
		for _, k := range Held.Nachteile {
			if v == k.Name {
				return
			}
		}
	}
	return true
}
