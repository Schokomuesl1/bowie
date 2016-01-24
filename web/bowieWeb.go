package bowieweb

import (
	"fmt"
	"github.com/Schokomuesl1/bowie/basiswerte"
	"github.com/Schokomuesl1/bowie/erschaffung"
	"github.com/Schokomuesl1/bowie/held"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type PageDataType struct {
	AlleEigenschaften *map[string]string
	AlleSpezies       *map[string]basiswerte.SpeziesType
	AlleTalente       *[]basiswerte.TalentType
	AlleKulturen      *map[string]basiswerte.KulturType
	AlleLiturgien     *map[string]basiswerte.LiturgieType
	Kosten            *map[string][26]int
	Grade             *map[string]erschaffung.Erfahrungsgrad
	Held              *held.Held
	Validator         *erschaffung.ErschaffungsValidator
	ValidatorMsg      []erschaffung.ValidatorMessage
}

type EigenschaftenModSet struct {
	Label        int
	Modifikation basiswerte.EigenschaftenModSpezies
}

var PageData PageDataType

func initPageData() {
	PageData = PageDataType{AlleEigenschaften: &basiswerte.AlleEigenschaften,
		AlleSpezies:   &basiswerte.AlleSpezies,
		AlleTalente:   &basiswerte.AlleTalente,
		AlleKulturen:  &basiswerte.AlleKulturen,
		AlleLiturgien: &basiswerte.AlleLiturgien,
		Kosten:        &basiswerte.Kostentable,
		Grade:         &erschaffung.AlleErfahrungsgrade,
		Held:          nil,
		Validator:     nil}
}

func renderTemplate(w http.ResponseWriter, tmpl string, pd *PageDataType) {
	t, _ := template.ParseFiles("template/" + tmpl + ".tpl")
	t.Execute(w, pd)
}

func resetHero(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: resetHero")
}

func addToValue(c web.C, w http.ResponseWriter, r *http.Request, addTo []string, val int) {
	if len(addTo) != 2 {
		return
	}
	group := addTo[0]
	item := addTo[1]
	switch group {
	case "eigenschaft":
		{
			e := PageData.Held.Eigenschaften.Get(item)
			if e != nil {
				if basiswerte.Kosten("E", e.Wert+val) > -1 {
					kosten := basiswerte.Kosten("E", e.Wert+val)
					if val < 0 {
						kosten *= -1
					}
					e.Add(val)
					PageData.Held.AP_spent += kosten
					PageData.Held.AP -= kosten
					fmt.Println(item, kosten, PageData.Held.AP_spent, PageData.Held.AP)
				}

			}
		}
	// if we have not found it in eigenschaften, it might be a Talent...
	case "talent":
		{
			t := PageData.Held.Talente.Get(item)
			if t != nil {
				if basiswerte.Kosten(t.SK, t.Value()+val) > -1 {
					kosten := basiswerte.Kosten(t.SK, t.Value()+val)
					if val < 0 {
						kosten *= -1
					}
					t.AddValue(val)
					PageData.Held.AP_spent += kosten
					PageData.Held.AP -= kosten
				}
			}
		}
	case "kampftechnik":
		{
			t := PageData.Held.Kampftechniken.Get(item)
			if t != nil {
				if basiswerte.Kosten(t.SK, t.Value()+val) > -1 {
					kosten := basiswerte.Kosten(t.SK, t.Value()+val)
					if val < 0 {
						kosten *= -1
					}
					t.AddValue(val)
					PageData.Held.AP_spent += kosten
					PageData.Held.AP -= kosten
				}
			}
		}
	}
}

func runActionParams(c web.C, w http.ResponseWriter, r *http.Request, action string, params []string) {
	switch action {
	case "increment":
		{
			addToValue(c, w, r, params, 1)
		}
	case "decrement":
		{
			addToValue(c, w, r, params, -1)
		}
	}
}

func runAction(c web.C, w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("Received request: runAction with parameters: action: %s, rest: %s", c.URLParams["action"], c.URLParams["*"])
	action := c.URLParams["action"]
	params := strings.Split(c.URLParams["*"], "/")
	if len(params) > 1 {
		runActionParams(c, w, r, action, params[1:])
	}
	if PageData.Validator != nil {
		_, PageData.ValidatorMsg = PageData.Validator.Validate()
	}
	renderTemplate(w, "held", &PageData)
}

func isValid(c web.C, w http.ResponseWriter, r *http.Request) {
}

func startPage(c web.C, w http.ResponseWriter, r *http.Request) {
	initPageData()
	renderTemplate(w, "start", &PageData)
}

func newHero(c web.C, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range PageData.Held.Spezies.EigenschaftsModifikationen {
		result := r.FormValue(strconv.Itoa(i))
		eigenschaft := PageData.Held.Eigenschaften.Get(result)
		if eigenschaft != nil {
			fmt.Println("start", eigenschaft)
			eigenschaft.SetMin(eigenschaft.Min() + v.Mod)
			eigenschaft.SetMax(eigenschaft.Max() + v.Mod)
			PageData.Held.Eigenschaften.Set(result, eigenschaft.Wert+v.Mod)
			fmt.Println("stop", eigenschaft)
		}
	}
	renderTemplate(w, "held", &PageData)
}

func modEigenschaften(c web.C, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	PageData.Held, PageData.Validator = erschaffung.ErschaffeHeld(r.FormValue("erfahrungsgrad"))
	PageData.Validator.AddAllValidators()
	PageData.Held.Name = r.FormValue("heldName")
	PageData.Held.SetSpezies(r.FormValue("spezies"))
	PageData.Held.SetKultur(r.FormValue("kultur"))
	PageData.Held.Eigenschaften.Init("MU", PageData.Validator.Grad.Eigenschaft)
	PageData.Held.Eigenschaften.Init("KL", PageData.Validator.Grad.Eigenschaft)
	PageData.Held.Eigenschaften.Init("GE", PageData.Validator.Grad.Eigenschaft)
	PageData.Held.Eigenschaften.Init("KK", PageData.Validator.Grad.Eigenschaft)
	PageData.Held.Eigenschaften.Init("FF", PageData.Validator.Grad.Eigenschaft)
	PageData.Held.Eigenschaften.Init("IN", PageData.Validator.Grad.Eigenschaft)
	PageData.Held.Eigenschaften.Init("CH", PageData.Validator.Grad.Eigenschaft)
	PageData.Held.Eigenschaften.Init("KO", PageData.Validator.Grad.Eigenschaft)
	PageData.Held.Talente.SetErschaffungsMax(PageData.Validator.Grad.Fertigkeit)
	_, PageData.ValidatorMsg = PageData.Validator.Validate()

	t, _ := template.ParseFiles("template/modSpeziesEigenschaften.tpl")

	eigenMod := make([]EigenschaftenModSet, len(PageData.Held.Spezies.EigenschaftsModifikationen))
	for i, v := range PageData.Held.Spezies.EigenschaftsModifikationen {
		eigenMod[i].Label = i
		eigenMod[i].Modifikation = v

	}
	t.Execute(w, &eigenMod)
}

func initRoutes() {
	// prepare routes, get/post stuff etc
	goji.Get("/", startPage)
	goji.Post("/held/reset", resetHero)
	goji.Post("/held/new", newHero)
	goji.Post("/held/modEigenschaften", modEigenschaften)
	goji.Get("/held/action/:action/*", runAction)
	goji.Get("/held/isValid", isValid)
}

func Serve() {
	initRoutes()
	goji.Serve()
}
