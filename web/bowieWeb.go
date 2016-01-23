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
	//"net/http/httputil"
	"strconv"
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

func incrementValue(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: incrementValue with parameters: name: %s", c.URLParams["name"])
}

func decrementValue(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: decrementValue with parameters: name: %s", c.URLParams["name"])
}

func canIncrementValue(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: canIncrementValue with parameters: name: %s", c.URLParams["name"])
	fmt.Fprintf(w, "{\"CanIncrement\":\"true\", \"Name\":\"%s\"}", c.URLParams["name"])
}

func canDecrementValue(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: canDecrementValue with parameters: name: %s", c.URLParams["name"])
}

func setValue(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: setValue with parameters: name: %s and tail %s.", c.URLParams["name"], c.URLParams["*"])
}

func getValue(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: getValue with parameters: name: %s and tail %s", c.URLParams["name"], c.URLParams["*"])
}

func isValid(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: isValid with no parameters")
	fmt.Fprintf(w, "{\"IsValid\":\"false\", \"Msg\":[\"Too many EP spent!\", \"Selbstbeherrschung too high!\"]}")
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
			PageData.Held.Eigenschaften.Set(result, eigenschaft.Wert+v.Mod)
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
	PageData.Validator.AddValidator(erschaffung.EPValidator{})
	PageData.Validator.AddValidator(erschaffung.FertigkeitsValidator{})
	PageData.Held.Name = r.FormValue("heldName")
	PageData.Held.SetSpezies(r.FormValue("spezies"))
	PageData.Held.SetKultur(r.FormValue("kultur"))
	PageData.Held.Eigenschaften.Set("MU", 8)
	PageData.Held.Eigenschaften.Set("KL", 8)
	PageData.Held.Eigenschaften.Set("GE", 8)
	PageData.Held.Eigenschaften.Set("KK", 8)
	PageData.Held.Eigenschaften.Set("FF", 8)
	PageData.Held.Eigenschaften.Set("IN", 8)
	PageData.Held.Eigenschaften.Set("CH", 8)
	PageData.Held.Eigenschaften.Set("KO", 8)
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
	//e.g. goji.Post("/held/reset", resetHero) // initial state (select name/species/culture...)
	goji.Get("/", startPage)
	goji.Post("/held/reset", resetHero)
	goji.Post("/held/new", newHero)
	goji.Post("/held/modEigenschaften", modEigenschaften)
	goji.Post("/held/increment/:name", incrementValue)
	goji.Post("/held/decrement/:name", decrementValue)
	goji.Get("/held/canincrement/:name", canIncrementValue) // true if enough AP available
	goji.Get("/held/candecrement/:name", canDecrementValue) // true if not at min value
	goji.Post("/held/set/:value/*", setValue)
	goji.Get("/held/get/:value/*", getValue)
	goji.Get("/held/isValid", isValid)
}

func Serve() {
	initRoutes()
	goji.Serve()
}
