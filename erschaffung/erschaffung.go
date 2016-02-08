package erschaffung

import (
	"fmt"
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
	h.ZauberCountMax = e.Grad.Zauber
	return h, e
}

func SFAvailable(Held *held.Held, SF *basiswerte.Sonderfertigkeit) (result bool, message string) {
	result = false
	message = ""
	// talente checken
	for _, v := range SF.Vorraussetzungen.Talente {
		t := Held.Talente.Get(v[0])
		v_num, err := strconv.Atoi(v[1])
		if t == nil || err != nil {
			return
		}
		if t.Value() < v_num {
			message = fmt.Sprintf("Anforderungen für Sonderfertigkeit %s nicht erfüllt! Talent %s muss mindestens %d betragen, aktuell ist der Wert jedoch %d", SF.Name, v[0], v_num, t.Value())
			return
		}
	}

	// Eigenschaften checken
	for _, v := range SF.Vorraussetzungen.Eigenschaften {
		t := Held.Eigenschaften.Get(v[0])
		v_num, err := strconv.Atoi(v[1])
		if t == nil || err != nil {
			return
		}
		if t.Value() < v_num {
			message = fmt.Sprintf("Anforderungen für Sonderfertigkeit %s nicht erfüllt! Eigenschaft %s muss mindestens %d betragen, aktuell ist der Wert jedoch %d", SF.Name, v[0], v_num, t.Value())
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
			message = fmt.Sprintf("Anforderungen für Sonderfertigkeit %s nicht erfüllt! Vorteil %s muss gewählt sein, fehlt jedoch.", SF.Name, v)
			return
		}
	}
	// Verbotene Nachteile...
	for _, v := range SF.Vorraussetzungen.NichtNachteil {
		for _, k := range Held.Vorteile {
			if k.Name == v {
				message = fmt.Sprintf("Anforderungen für Sonderfertigkeit %s nicht erfüllt! Nachteil %s darf nicht gewählt sein!", SF.Name, v)
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
		for _, k := range Held.Sonderfertigkeiten.Karmale {
			if k.Name == v {
				found = true
				break
			}
		}

		for _, k := range Held.Sonderfertigkeiten.Magische {
			if k.Name == v {
				found = true
				break
			}
		}

		for _, k := range Held.Sonderfertigkeiten.Kampf {
			if k.Name == v {
				found = true
				break
			}
		}

		for _, k := range Held.Sonderfertigkeiten.Schriften {
			if k.Name == v {
				found = true
				break
			}
		}

		for _, k := range Held.Sonderfertigkeiten.Sprachen {
			if k.Name == v {
				found = true
				break
			}
		}

		if !found {
			message = fmt.Sprintf("Anforderungen für Sonderfertigkeit %s nicht erfüllt! SF %s muss gewählt sein, fehlt jedoch.", SF.Name, v)
			return
		}
	}

	return true, ""
}

func VorUndNachteilAvailable(Held *held.Held, VorOderNachteil *basiswerte.VorUndNachteil) (result bool, message string) {
	result = false
	message = ""
	for _, v := range VorOderNachteil.Vorraussetzungen.Vorteile {
		found := false
		for _, k := range Held.Vorteile {
			if v == k.Name {
				found = true
				break
			}
		}
		if !found {
			message = fmt.Sprintf("Anforderungen für Vorteil/Nachteil %s nicht erfüllt! Vorteil %s muss gewählt sein, fehlt jedoch.", VorOderNachteil.Name, v)
			return
		}
	}

	for _, v := range VorOderNachteil.Vorraussetzungen.Nachteile {
		found := false
		for _, k := range Held.Nachteile {
			if v == k.Name {
				found = true
				message = fmt.Sprintf("Anforderungen für Vorteil/Nachteil %s nicht erfüllt! Nachteil %s muss gewählt sein, fehlt jedoch.", VorOderNachteil.Name, v)
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
				message = fmt.Sprintf("Anforderungen für Vorteil/Nachteil %s nicht erfüllt! Vorteil %s darf nicht gewählt sein, ist jedoch vorhanden.", VorOderNachteil.Name, v)
				return
			}
		}
	}

	for _, v := range VorOderNachteil.Vorraussetzungen.NichtNachteile {
		for _, k := range Held.Nachteile {
			if v == k.Name {
				message = fmt.Sprintf("Anforderungen für Vorteil/Nachteil %s nicht erfüllt! Nachteil %s darf nicht gewählt sein, ist jedoch vorhanden.", VorOderNachteil.Name, v)
				return
			}
		}
	}
	return true, ""
}
