package erschaffung

import "github.com/Schokomuesl1/bowie/held"

func ErschaffeHeld(erfahrungsgrad string) (*held.Held, *ErschaffungsValidator) {
	h := held.NewHeld()
	e := MakeErschaffungsValidator(h, erfahrungsgrad)
	if e == nil {
		return nil, nil
	}
	h.AP = e.Grad.AP
	return h, e
}
