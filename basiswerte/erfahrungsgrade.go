package basiswerte

import (
	"encoding/json"
	"io/ioutil"
)

var AlleErfahrungsgrade map[string]Erfahrungsgrad

func init() {
	file, _ := ioutil.ReadFile("regeln/erfahrungsgrade.json")
	erfahrungsgradTmp := make([]Erfahrungsgrad, 0)
	json.Unmarshal([]byte(string(file)), &erfahrungsgradTmp)
	AlleErfahrungsgrade = make(map[string]Erfahrungsgrad)
	for _, v := range erfahrungsgradTmp {
		AlleErfahrungsgrade[v.Name] = v
	}
}

type Erfahrungsgrad struct {
	Name         string
	AP           int
	Eigenschaft  int
	Fertigkeit   int
	Kampftechnik int
	EP           int
	Zauber       [2]int
}
