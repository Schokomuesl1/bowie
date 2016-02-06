package erschaffung

import (
	"fmt"
	"github.com/Schokomuesl1/bowie/held"
)

type ValidatorMessageType int

const (
	NONE ValidatorMessageType = iota
	INFO
	WARNING
	ERROR
)

func (e ValidatorMessageType) String() string {
	switch e {
	case NONE:
		return ""
	case INFO:
		return "Info"
	case WARNING:
		return "Warnung"
	case ERROR:
		return "Fehler"
	default:
		return "Falscher Wert! Programmfehler!"
	}
}

type Validator interface {
	Validate(*Erfahrungsgrad, *held.Held) (bool, ValidatorMessage)
}

type ValidatorMessage struct {
	Msg  string
	Type ValidatorMessageType
}

type ErschaffungsValidator struct {
	Grad        Erfahrungsgrad
	Held        *held.Held
	Validatoren []Validator
}

func MakeErschaffungsValidator(held *held.Held, erfahrungsgrad string) *ErschaffungsValidator {
	_, exists := AlleErfahrungsgrade[erfahrungsgrad]
	if !exists {
		return nil
	}
	if held == nil {
		return nil
	}
	return &ErschaffungsValidator{
		Grad:        AlleErfahrungsgrade[erfahrungsgrad],
		Held:        held,
		Validatoren: make([]Validator, 0)}
}

func (e *ErschaffungsValidator) Validate() (result bool, messages []ValidatorMessage) {
	result = true
	messages = make([]ValidatorMessage, 0, 1)
	for _, v := range e.Validatoren {
		validatorResult, message := v.Validate(&e.Grad, e.Held)
		result = result && validatorResult
		if len(message.Msg) != 0 {
			n := len(messages)
			// possibly extend slice/buffer
			if n == cap(messages) {
				newMessages := make([]ValidatorMessage, len(messages), 2*len(messages)+1)
				copy(newMessages, messages)
				messages = newMessages
			}
			messages = messages[0 : n+1]
			messages[n] = message
		}
	}
	n := len(messages)
	if n == cap(messages) {
		newMessages := make([]ValidatorMessage, len(messages), len(messages)+1) // we only need one more
		copy(newMessages, messages)
		messages = newMessages
	}
	if result {
		messages = messages[0 : n+1]
		messages[n] = ValidatorMessage{Msg: fmt.Sprintf("Held %s ist valide!", e.Held.Name), Type: INFO}
	}
	return
}

func (e *ErschaffungsValidator) AddValidator(v Validator) {
	n := len(e.Validatoren)
	// possibly extend slice/buffer
	if n == cap(e.Validatoren) {
		newValidators := make([]Validator, len(e.Validatoren), 2*len(e.Validatoren)+1)
		copy(newValidators, e.Validatoren)
		e.Validatoren = newValidators
	}
	e.Validatoren = e.Validatoren[0 : n+1]
	e.Validatoren[n] = v
}

func (e *ErschaffungsValidator) AddAllValidators() {
	e.AddValidator(EPValidator{})
	e.AddValidator(FertigkeitsValidator{})
	e.AddValidator(APValidator{})
	e.AddValidator(ZauberUndLiturgieValidator{})
	e.AddValidator(VorteilUndNachteilValidator{vorteil: true})
	e.AddValidator(VorteilUndNachteilValidator{vorteil: false})
}

// validators here

// validate max. Eigenschaftspunkte
type EPValidator struct {
}

func (e EPValidator) Validate(grad *Erfahrungsgrad, held *held.Held) (result bool, message ValidatorMessage) {
	result = true
	message.Msg = ""
	message.Type = NONE
	sum := 0
	for _, v := range held.Eigenschaften.Eigenschaften {
		sum += v.Value()
	}
	result = sum <= grad.EP
	if !result {
		message.Msg = fmt.Sprintf("Zu viele EP verbraucht: %d von %d verfügbaren!", sum, grad.EP)
		message.Type = ERROR
	} else {
		message.Msg = fmt.Sprintf("Aktuell verbraucht: %d von %d verfügbaren.", sum, grad.EP)
		message.Type = INFO
	}
	return
}

// validate max. Fertigkeitslevel
type FertigkeitsValidator struct {
}

func (e FertigkeitsValidator) Validate(grad *Erfahrungsgrad, held *held.Held) (result bool, message ValidatorMessage) {
	result = true
	message.Msg = ""
	message.Type = NONE

	for _, v := range held.Zauber.Zaubers {
		if v.Value() > 14 {
			result = false
			message.Type = ERROR
			message.Msg = fmt.Sprintf("Zauber %s hat einen Wert von %d. Maximum bei Erschaffung ist 14! ", v.Name, v.Value())
			return
		}
	}
	for _, v := range held.Liturgien.Liturgien {
		if v.Value() > 14 {
			result = false
			message.Type = ERROR
			message.Msg = fmt.Sprintf("Liturgie %s hat einen Wert von %d. Maximum bei Erschaffung ist 14! ", v.Name, v.Value())
			return
		}
	}

	message.Type = INFO
	message.Msg = fmt.Sprintf("Keine Zauber und Liturgien mit einem Wert größer 14 gefunden!")
	return
}

// validate max. Zauber und Liturgie-Grad
type ZauberUndLiturgieValidator struct {
}

func (e ZauberUndLiturgieValidator) Validate(grad *Erfahrungsgrad, held *held.Held) (result bool, message ValidatorMessage) {
	result = true
	message.Msg = ""
	message.Type = NONE

	for _, v := range held.Talente.Talente {
		if v.Value() > grad.Fertigkeit {
			result = false
			message.Type = ERROR
			message.Msg = fmt.Sprintf("Talent %s hat einen Wert von %d. Maximum für Erfahrungsstufe %s ist %d.", v.Name, v.Value(), grad.Name, grad.Fertigkeit)
			return
		}
	}
	message.Type = INFO
	message.Msg = fmt.Sprintf("Keine Fertigkeit mit einem Wert größer %d gefunden!", grad.Fertigkeit)
	return
}

// Validator ausgegebene AP
type VorteilUndNachteilValidator struct {
	vorteil bool
}

func (e VorteilUndNachteilValidator) Validate(grad *Erfahrungsgrad, held *held.Held) (result bool, message ValidatorMessage) {
	result = true
	message.Msg = ""
	message.Type = NONE
	sum := 0
	if e.vorteil {
		for _, v := range held.Vorteile {
			sum += v.APKosten
		}
		if sum > 80 {
			message.Msg = fmt.Sprintf("Zuviele Vorteile! Maximal 80 AP in Vorteilen - genutzt: %d!", sum)
			message.Type = ERROR
		}

	} else {
		for _, v := range held.Nachteile {
			sum += v.APKosten
		}
		if sum < -80 {
			message.Msg = fmt.Sprintf("Zuviele Nachteile! Maximal -80 AP in Nachteilen - genutzt: %d!", sum)
			message.Type = ERROR
		}
	}
	return
}

// Validator ausgegebene AP
type APValidator struct {
}

func (e APValidator) Validate(grad *Erfahrungsgrad, held *held.Held) (result bool, message ValidatorMessage) {
	result = true
	message.Msg = ""
	message.Type = NONE

	if held.AP < 0 {
		result = false
		message.Type = ERROR
		message.Msg = fmt.Sprintf("%d AP zuviel verbraucht!", held.AP)
	}
	return
}
