package main

import (
	//"encoding/json"
	"fmt"
	"github.com/Schokomuesl1/bowie/erschaffung"
	//"github.com/Schokomuesl1/bowie/held"

	//"io/ioutil"
)

type EigenschaftEntry struct {
	Kurz string
	Lang string
}

func main() {
	//KL := held.MakeEigenschaft("KL", 12)
	/*MU := held.MakeEigenschaft("MU", 12)
	    GE := held.MakeEigenschaft("GE", 12)
	    KK := held.MakeEigenschaft("KK", 12)
	    //FF := held.MakeEigenschaft("FF", 12)
	    //IN := held.MakeEigenschaft("IN", 12)
	    //CH := held.MakeEigenschaft("CH", 12)
	    t := held.MakeTalent("Selbstbeherrschung", 5, MU, KK, GE)
	    fmt.Println(t)
	    MU.Increment()
	    fmt.Println(t)
	    ini := held.MakeCalculatedDependentValue("Initiative", 0.5, []held.DependentValue{MU, GE})
	    fmt.Println(ini)
	    success, ergebnis := KK.Probe()

	    for i := -5; i < 6; i++ {
	  	success, ergebnis := KK.ProbeMod(i)
	  	fmt.Println(success, ergebnis)
	    }
	    fmt.Println("Probe: KK: ", KK.Value(), success, ergebnis)
	    fmt.Println(t.Probe())*/
	/*h := held.NewHeld()
	validator := erschaffung.MakeErschaffungsValidator(h, "Kompetent")*/
	h, validator := erschaffung.ErschaffeHeld("Kompetent")
	validator.AddValidator(erschaffung.EPValidator{})
	validator.AddValidator(erschaffung.FertigkeitsValidator{})
	h.Eigenschaften.Set("MU", 8)
	h.Eigenschaften.Set("KL", 9)
	h.Eigenschaften.Set("GE", 10)
	h.Eigenschaften.Set("KK", 11)
	h.Eigenschaften.Set("FF", 12)
	h.Eigenschaften.Set("IN", 13)
	h.Eigenschaften.Set("CH", 14)
	h.Eigenschaften.Set("KO", 15)
	h.SetSpezies("Mensch")
	h.SetKultur("Aranier")
	h.Eigenschaften.Set("GE", 15)
	fmt.Println(h)
	result, messages := validator.Validate()
	fmt.Println(result)
	for _, v := range messages {
		fmt.Println(v)
	}
	h.Talente.Get("Verbergen").Wert = 20
	h.Eigenschaften.Set("MU", 13)
	h.Eigenschaften.Set("KL", 15)
	h.Eigenschaften.Set("GE", 15)
	result, messages = validator.Validate()
	fmt.Println(result)
	for _, v := range messages {
		fmt.Println(v)
	}
}
