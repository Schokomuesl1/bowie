package main

import (
	//"encoding/json"
	"fmt"
	"github.com/Schokomuesl1/bowie/erschaffung"
	"github.com/Schokomuesl1/bowie/held"
	//"github.com/Schokomuesl1/bowie/web"
	"html/template"
	"net/http"
	//"os"
	//"github.com/Schokomuesl1/bowie/held"
	"github.com/Schokomuesl1/bowie/basiswerte"

	//"io/ioutil"
)

type EigenschaftEntry struct {
	Kurz string
	Lang string
}

type PageData struct {
	Held *held.Held
	Msg  []erschaffung.ValidatorMessage
}

func MakeHeld() *PageData {
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
	_, messages = validator.Validate()
	return &PageData{Held: h, Msg: messages}
}

func renderTemplate(w http.ResponseWriter, tmpl string, pd *PageData) {
	t, _ := template.ParseFiles("template/" + tmpl + ".tpl")
	t.Execute(w, pd)
}

func heldHandler(w http.ResponseWriter, r *http.Request) {
	pd := MakeHeld()
	renderTemplate(w, "held", pd)
}

func main() {
	//bowieweb.Serve()
	//http.HandleFunc("/held/", heldHandler)
	//http.ListenAndServe(":8080", nil)
	for _, v := range basiswerte.AlleLiturgien {
		fmt.Println(v)
	}
}
