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
	//	"net/http/httputil"
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
	Available         AvailableItems
}

type AvailableItems struct {
	Nachteile    []basiswerte.VorUndNachteil
	Vorteile     []basiswerte.VorUndNachteil
	SF_Allgemein []basiswerte.Sonderfertigkeit
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

func calculateAvailable() {

	PageData.Available.Nachteile = PageData.Available.Nachteile[:0]
	PageData.Available.Vorteile = PageData.Available.Vorteile[:0]
	PageData.Available.SF_Allgemein = PageData.Available.SF_Allgemein[:0]
	for _, v := range basiswerte.Nachteile {
		if erschaffung.VorUndNachteilAvailable(PageData.Held, &v) {
			// only append if not already selected
			selected := false
			for _, w := range PageData.Held.Nachteile {
				if w.Name == v.Name {
					selected = true
					break
				}
			}
			if !selected {
				PageData.Available.Nachteile = append(PageData.Available.Nachteile, v)
			}
		}
	}
	for _, v := range basiswerte.Vorteile {
		if erschaffung.VorUndNachteilAvailable(PageData.Held, &v) {
			// only append if not already selected
			selected := false
			for _, w := range PageData.Held.Vorteile {
				if w.Name == v.Name {
					selected = true
					break
				}
			}
			if !selected {
				PageData.Available.Vorteile = append(PageData.Available.Vorteile, v)
			}
		}
	}
	for _, v := range basiswerte.AllgemeineSF {
		if erschaffung.SFAvailable(PageData.Held, &v) {
			// only append if not already selected
			selected := false
			for _, w := range PageData.Held.Sonderfertigkeiten {
				if w.Name == v.Name {
					selected = true
					break
				}
			}
			if !selected {
				PageData.Available.SF_Allgemein = append(PageData.Available.SF_Allgemein, v)
			}
		}
	}
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

func addItem(c web.C, w http.ResponseWriter, r *http.Request, addTo []string) {
	if len(addTo) != 2 {
		return
	}
	group := addTo[0]
	item := addTo[1]
	switch group {
	case "vorteil":
		{
			vorteil := basiswerte.GetVorteil(item)
			if vorteil != nil {
				for _, v := range PageData.Held.Vorteile {
					if v.Name == vorteil.Name {
						return
					}
				}
				PageData.Held.Vorteile = append(PageData.Held.Vorteile, vorteil)
			}
		}
	case "nachteil":
		{
			nachteil := basiswerte.GetNachteil(item)
			if nachteil != nil {
				for _, v := range PageData.Held.Nachteile {
					if v.Name == nachteil.Name {
						return
					}
				}
				PageData.Held.Nachteile = append(PageData.Held.Nachteile, nachteil)
			}
		}
	case "sf":
		{
			sf := basiswerte.GetSF(item)
			if sf != nil {
				for _, v := range PageData.Held.Sonderfertigkeiten {
					if v.Name == sf.Name {
						return
					}
				}
				PageData.Held.Sonderfertigkeiten = append(PageData.Held.Sonderfertigkeiten, sf)
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
	case "add":
		{
			addItem(c, w, r, params)
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

	calculateAvailable()
	renderTemplate(w, "held", &PageData)
}

/*func addSF(c web.C, w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	defer calculateAvailable()
	renderTemplate(w, "held", &PageData)

}

func addVN(c web.C, w http.ResponseWriter, r *http.Request) {
	a, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(a[:]))
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	defer calculateAvailable()
	renderTemplate(w, "held", &PageData)

	result := r.FormValue("VorteilToAdd")
	if len(result) > 0 {
		vorteil := basiswerte.GetVorteil(result)
		if vorteil != nil {
			PageData.Held.Vorteile = append(PageData.Held.Vorteile, vorteil)
		}

		return
	}

	result = r.FormValue("NachteilToAdd")
	if len(result) > 0 {
		nachteil := basiswerte.GetNachteil(result)
		if nachteil != nil {
			PageData.Held.Nachteile = append(PageData.Held.Nachteile, nachteil)
		}
		return
	}
}
*/
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
	fmt.Println("parse ok")
	for i, v := range PageData.Held.Spezies.EigenschaftsModifikationen {
		fmt.Println(v)
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

	calculateAvailable()
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
	PageData.Held.APAusgeben(basiswerte.AlleSpezies[r.FormValue("spezies")].APKosten)
	PageData.Held.SetKultur(r.FormValue("kultur"))
	PageData.Held.APAusgeben(basiswerte.AlleKulturen[r.FormValue("kultur")].APKosten)
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
	/*goji.Post("/held/addSF", addSF)
	goji.Post("/held/addVN", addVN)*/
	goji.Get("/held/action/:action/*", runAction)
	goji.Get("/held/isValid", isValid)
}

func Serve() {
	initRoutes()
	goji.Serve()
}
