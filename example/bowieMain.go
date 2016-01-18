package main

import (
	"encoding/json"
	"fmt"
	"github.com/Schokomuesl1/bowie"
	"io/ioutil"
)

type EigenschaftEntry struct {
	Kurz string
	Lang string
}

func main() {
	//KL := bowie.MakeEigenschaft("KL", 12)
	/*MU := bowie.MakeEigenschaft("MU", 12)
	GE := bowie.MakeEigenschaft("GE", 12)
	KK := bowie.MakeEigenschaft("KK", 12)
	//FF := bowie.MakeEigenschaft("FF", 12)
	//IN := bowie.MakeEigenschaft("IN", 12)
	//CH := bowie.MakeEigenschaft("CH", 12)
	t := bowie.MakeTalent("Selbstbeherrschung", 5, MU, KK, GE)
	fmt.Println(t)
	MU.Increment()
	fmt.Println(t)
	ini := bowie.MakeCalculatedDependentValue("Initiative", 0.5, []bowie.DependentValue{MU, GE})
	fmt.Println(ini)
	success, ergebnis := KK.Probe()

	for i := -5; i < 6; i++ {
		success, ergebnis := KK.ProbeMod(i)
		fmt.Println(success, ergebnis)
	}
	fmt.Println("Probe: KK: ", KK.Value(), success, ergebnis)
	fmt.Println(t.Probe())*/
	h := bowie.NewHeld()
	h.Eigenschaften.Set("MU", 8)
	//h.Eigenschaften.Set("KL", 9)
	//h.Eigenschaften.Set("GE", 10)
	//h.Eigenschaften.Set("KK", 11)
	//h.Eigenschaften.Set("FF", 12)
	//h.Eigenschaften.Set("IN", 13)
	h.Eigenschaften.Set("CH", 14)
	//h.Eigenschaften.Set("KO", 15)
	//h.NewTalent("Selbstbeherrschung", [3]string{"MU", "KK", "KO"})
	fmt.Print(h)
	h.Eigenschaften.Set("KO", 15)
	fmt.Print(h) /*file, _ := ioutil.ReadFile("data/eigenschaften.json")
	//tmp := make([]EigenschaftEntry, 0)
	tmp := make(map[string]string)
	json.Unmarshal([]byte(string(file)), &tmp)
	fmt.Print(tmp)*/
	file, _ := ioutil.ReadFile("data/talente.json")
	AlleTalente := make([]bowie.TalentParse, 0)
	json.Unmarshal([]byte(string(file)), &AlleTalente)
	for _, v := range AlleTalente {
		fmt.Println(v)
		h.NewTalent(v.Name, v.Probe)
	}
	fmt.Print(h)
}
