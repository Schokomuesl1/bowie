package erschaffung

import (
	"fmt"
	//	"github.com/Schokomuesl1/bowie/basiswerte"
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
	Validate(*held.Held) (bool, ValidatorMessage)
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
	fmt.Println("before loop")
	fmt.Println(e.Validatoren)
	for _, v := range e.Validatoren {
		validatorResult, message := v.Validate(e.Held)
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
	fmt.Println("after loop")
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

// validators here

// validate max. Eigenschaftspunkte
type EPValidator struct {
	maxEP int
}

func (e *EPValidator) Validate(held *held.Held) (result bool, message ValidatorMessage) {
	result = true
	message.Msg = ""
	message.Type = NONE
	sum := 0
	for _, v := range held.Eigenschaften.Eigenschaften {
		sum += v.Value()
	}
	result = sum <= e.maxEP
	if !result {
		message.Msg = fmt.Sprintf("Zu viele EP verbraucht: %d von %d verfügbaren!", sum, e.maxEP)
	}
	return

}
