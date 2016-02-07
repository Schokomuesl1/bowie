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
	"encoding/json"
	"strconv"
	"strings"
)

type PageDataType struct {
	AlleEigenschaften *map[string]string
	AlleSpezies       *map[string]basiswerte.SpeziesType
	AlleTalente       *[]basiswerte.TalentType
	AlleKulturen      *map[string]basiswerte.KulturType
	AlleLiturgien     *map[string]basiswerte.LiturgieType
	AlleZauber        *map[string]basiswerte.ZauberType
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
	SF_Karmal    []basiswerte.Sonderfertigkeit
	SF_Magisch   []basiswerte.Sonderfertigkeit
	SF_Kampf     []basiswerte.Sonderfertigkeit
}

type EigenschaftenModSet struct {
	Label        int
	Modifikation basiswerte.EigenschaftenModSpezies
}

type apData struct {
	AP       int `json:"ap"`
	AP_spent int `json:"ap_spent"`
	AP_total int `json:"ap_total"`
}

func (a *apData) ProzentVerfuegbar() int {
	return a.AP / (a.AP_total / 100)
}

type redirectToStruct struct {
	RedirectTo        string                         `json:"redirectTo"`
	NotificationMsg   string                         `json:"notificationMsg"`
	Magie             bool                           `json:"magie"`
	Karmal            bool                           `json:"karmal"`
	ValidatorMessages []erschaffung.ValidatorMessage `json:"validatorMessages"`
}

var PageData PageDataType

func initPageData() {
	PageData = PageDataType{AlleEigenschaften: &basiswerte.AlleEigenschaften,
		AlleSpezies:   &basiswerte.AlleSpezies,
		AlleTalente:   &basiswerte.AlleTalente,
		AlleKulturen:  &basiswerte.AlleKulturen,
		AlleLiturgien: &basiswerte.AlleLiturgien,
		AlleZauber:    &basiswerte.AlleZauber,
		Kosten:        &basiswerte.Kostentable,
		Grade:         &erschaffung.AlleErfahrungsgrade,
		Held:          nil,
		Validator:     nil}
}

func calculateAvailable() {

	PageData.Available.Nachteile = PageData.Available.Nachteile[:0]
	PageData.Available.Vorteile = PageData.Available.Vorteile[:0]
	PageData.Available.SF_Allgemein = PageData.Available.SF_Allgemein[:0]
	PageData.Available.SF_Karmal = PageData.Available.SF_Karmal[:0]
	PageData.Available.SF_Magisch = PageData.Available.SF_Magisch[:0]
	PageData.Available.SF_Kampf = PageData.Available.SF_Kampf[:0]
	for _, v := range basiswerte.Nachteile {
		ok, _ := erschaffung.VorUndNachteilAvailable(PageData.Held, &v)
		if ok {
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
		ok, _ := erschaffung.VorUndNachteilAvailable(PageData.Held, &v)
		if ok {
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
		ok, _ := erschaffung.SFAvailable(PageData.Held, &v)
		if ok {
			// only append if not already selected
			selected := false
			for _, w := range PageData.Held.Sonderfertigkeiten.Allgemeine {
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
	for _, v := range basiswerte.KarmaleSF {
		ok, _ := erschaffung.SFAvailable(PageData.Held, &v)
		if ok {
			// only append if not already selected
			selected := false
			for _, w := range PageData.Held.Sonderfertigkeiten.Karmale {
				if w.Name == v.Name {
					selected = true
					break
				}
			}
			if !selected {
				PageData.Available.SF_Karmal = append(PageData.Available.SF_Karmal, v)
			}
		}
	}
	for _, v := range basiswerte.MagischeSF {
		ok, _ := erschaffung.SFAvailable(PageData.Held, &v)
		if ok {
			// only append if not already selected
			selected := false
			for _, w := range PageData.Held.Sonderfertigkeiten.Magische {
				if w.Name == v.Name {
					selected = true
					break
				}
			}
			if !selected {
				PageData.Available.SF_Magisch = append(PageData.Available.SF_Magisch, v)
			}
		}
	}
	for _, v := range basiswerte.KampfSF {
		ok, _ := erschaffung.SFAvailable(PageData.Held, &v)
		if ok {
			// only append if not already selected
			selected := false
			for _, w := range PageData.Held.Sonderfertigkeiten.Kampf {
				if w.Name == v.Name {
					selected = true
					break
				}
			}
			if !selected {
				PageData.Available.SF_Kampf = append(PageData.Available.SF_Kampf, v)
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

func addToValue(c web.C, w http.ResponseWriter, r *http.Request, addTo []string, val int) (string, string) {
	if len(addTo) != 2 {
		return "", ""
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
					PageData.Held.APAusgeben(kosten)
					return "/held/page/allgemeines", ""
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
					PageData.Held.APAusgeben(kosten)
					return "/held/page/talente", ""
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
					PageData.Held.APAusgeben(kosten)
					return "/held/page/kampftechniken", ""
				}
			}
		}
	case "zauber":
		{
			z := PageData.Held.Zauber.Get(item)
			if z != nil {
				kosten := basiswerte.Kosten(z.SK(), z.Value()+val)
				if kosten > -1 {
					if val < 0 {
						kosten *= -1
					}
					z.AddValue(val)
					PageData.Held.APAusgeben(kosten)
					return "/held/page/magie", ""
				}
			}
		}
	case "liturgie":
		{
			l := PageData.Held.Liturgien.Get(item)
			if l != nil {
				kosten := basiswerte.Kosten(l.SK(), l.Value()+val)
				if kosten > -1 {
					if val < 0 {
						kosten *= -1
					}
					l.AddValue(val)
					PageData.Held.APAusgeben(kosten)
					return "/held/page/karmales", ""
				}
			}
		}
	}
	return "", ""
}

func removeItem(c web.C, w http.ResponseWriter, r *http.Request, addTo []string) (string, string) {
	if len(addTo) != 2 {
		return "", ""
	}
	//Vorteile                   []string
	//Nachteile                  []string
	group := addTo[0]
	item := addTo[1]
	switch group {
	case "vorteilnachteil":
		{
			vorteil := basiswerte.GetVorteil(item)
			if vorteil != nil {
				for i, v := range PageData.Held.Vorteile {
					if v.Name == vorteil.Name {
						// ok, we need to check if the vorteil stems from the species descrtipion (e.g. Elf -> Zauberer). If so-> no deleting
						for _, vt := range PageData.Held.Spezies.Vorteile {
							if vorteil.Name == vt {
								return "", ""
							}
						}
						PageData.Held.Vorteile = append(PageData.Held.Vorteile[:i], PageData.Held.Vorteile[i+1:]...)
						PageData.Held.APAusgeben(vorteil.APKosten * -1)
						return "/held/page/allgemeines", ""
					}
				}
			}
			// no vorteil found -> continue...
			nachteil := basiswerte.GetNachteil(item)
			if nachteil != nil {
				for i, v := range PageData.Held.Nachteile {
					if v.Name == nachteil.Name {
						// ok, we need to check if the vorteil stems from the species descrtipion (e.g. Elf -> Zauberer). If so-> no deleting
						for _, nt := range PageData.Held.Spezies.Nachteile {
							if nachteil.Name == nt {
								return "", ""
							}
						}
						PageData.Held.Nachteile = append(PageData.Held.Nachteile[:i], PageData.Held.Nachteile[i+1:]...)
						PageData.Held.APAusgeben(nachteil.APKosten * -1)
						return "/held/page/allgemeines", ""
					}
				}
			}
		}
	case "SFToAddAllgemein", "SFToAddKarmal", "SFToAddMagisch", "SFToAddKampf":
		{
			var bereich *[]*basiswerte.Sonderfertigkeit
			switch group {
			case "SFToAddAllgemein":
				bereich = &PageData.Held.Sonderfertigkeiten.Allgemeine
			case "SFToAddKarmal":
				bereich = &PageData.Held.Sonderfertigkeiten.Karmale
			case "SFToAddMagisch":
				bereich = &PageData.Held.Sonderfertigkeiten.Magische
			case "SFToAddKampf":
				bereich = &PageData.Held.Sonderfertigkeiten.Kampf
			default:
				return "", ""
			}
			sf := basiswerte.GetSF(item)
			if sf != nil {
				for i, v := range *bereich {
					if v.Name == sf.Name {
						(*bereich), (*bereich)[len((*bereich))-1] = append((*bereich)[:i], (*bereich)[i+1:]...), nil
						PageData.Held.APAusgeben(sf.APKosten * -1)
						switch group {
						case "SFToAddAllgemein":
							return "/held/page/allgemeines", ""
						case "SFToAddKarmal":
							return "/held/page/karmales", ""
						case "SFToAddMagisch":
							return "/held/page/magie", ""
						case "SFToAddKampf":
							return "/held/page/kampftechniken", ""
						default:
							return "", ""
						}
					}
				}
			}
		}
	}
	return "", ""
}

func addItem(c web.C, w http.ResponseWriter, r *http.Request, addTo []string) (string, string) {
	if len(addTo) != 2 {
		return "", ""
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
						return "", ""
					}
				}
				PageData.Held.Vorteile = append(PageData.Held.Vorteile, *vorteil)
				PageData.Held.APAusgeben(vorteil.APKosten)
				return "/held/page/allgemeines", ""
			}
		}
	case "nachteil":
		{
			nachteil := basiswerte.GetNachteil(item)
			if nachteil != nil {
				for _, v := range PageData.Held.Nachteile {
					if v.Name == nachteil.Name {
						return "", ""
					}
				}
				PageData.Held.Nachteile = append(PageData.Held.Nachteile, *nachteil)
				PageData.Held.APAusgeben(nachteil.APKosten)
				return "/held/page/allgemeines", ""
			}
		}
	case "SFToAddAllgemein", "SFToAddKarmal", "SFToAddMagisch", "SFToAddKampf":
		{
			var bereich *[]*basiswerte.Sonderfertigkeit
			switch group {
			case "SFToAddAllgemein":
				bereich = &PageData.Held.Sonderfertigkeiten.Allgemeine
			case "SFToAddKarmal":
				bereich = &PageData.Held.Sonderfertigkeiten.Karmale
			case "SFToAddMagisch":
				bereich = &PageData.Held.Sonderfertigkeiten.Magische
			case "SFToAddKampf":
				bereich = &PageData.Held.Sonderfertigkeiten.Kampf
			default:
				return "", ""
			}
			sf := basiswerte.GetSF(item)
			if sf != nil {
				for _, v := range *bereich {
					if v.Name == sf.Name {
						return "", ""
					}
				}
				*bereich = append(*bereich, sf)
				PageData.Held.APAusgeben(sf.APKosten)
				switch group {
				case "SFToAddAllgemein":
					return "/held/page/allgemeines", ""
				case "SFToAddKarmal":
					return "/held/page/karmales", ""
				case "SFToAddMagisch":
					return "/held/page/magie", ""
				case "SFToAddKampf":
					return "/held/page/kampftechniken", ""
				default:
					return "", ""
				}

			}
		}
	case "zauber":
		{

			ok, _ := erschaffung.VorUndNachteilAvailable(PageData.Held, basiswerte.GetVorteil("Zauberer"))
			if !ok {
				return "", ""
			}
			_, exists := basiswerte.AlleZauber[item]
			if !exists {
				return "", ""
			}
			zauber, _ := basiswerte.AlleZauber[item]
			// first we need to check if there are still zauber left to add in this category...
			ownCategory := false
			fmt.Println(zauber)
			for _, v := range zauber.Verbreitung {
				fmt.Println(v)
				if v == "allgemein" || v == "Allgemein" {
					ownCategory = true
					break
				}
				lookForSF := fmt.Sprintf("Tradition (%s)", v)
				for _, s := range PageData.Held.Sonderfertigkeiten.Magische {
					if s.Name == lookForSF {
						ownCategory = true
						break
					}
				}
			}
			idx := 0
			if !ownCategory {
				idx = 1
			}
			if PageData.Held.ZauberCount[idx] >= PageData.Validator.Grad.Zauber[idx] {
				// no room left, we cant add it...
				return "", "Kann Zauber nicht hinzufügen, maximale Anzahl in dieser Kategorie bereits erreicht."
			}
			if PageData.Held.Zauber.Exists(item) {
				return "", "Der Held beherrscht diesen Zauber bereits."
			}
			PageData.Held.ZauberCount[idx]++
			PageData.Held.NewZauber(&zauber)
			PageData.Held.Zauber.Get(item).SetMaxErschaffung(PageData.Validator.Grad.Fertigkeit)
			if zauber.Steigerungsfaktor != "-" {
				PageData.Held.APAusgeben(basiswerte.Kosten(zauber.Steigerungsfaktor, 0))
			} else {
				PageData.Held.APAusgeben(1) // Zaubertrick + Segnung 1 AP
			}
			return "/held/page/magie", ""
		}
	case "liturgie":
		{
			ok, _ := erschaffung.VorUndNachteilAvailable(PageData.Held, basiswerte.GetVorteil("Geweihter"))
			if !ok {
				return "", ""
			}
			_, exists := basiswerte.AlleLiturgien[item]
			if !exists {
				return "", ""
			}
			liturgie, _ := basiswerte.AlleLiturgien[item]
			PageData.Held.NewLiturgie(&liturgie)
			PageData.Held.Liturgien.Get(item).SetMaxErschaffung(PageData.Validator.Grad.Fertigkeit)
			if liturgie.Steigerungsfaktor != "-" {
				PageData.Held.APAusgeben(basiswerte.Kosten(liturgie.Steigerungsfaktor, 0))
			} else {
				PageData.Held.APAusgeben(1) // Zaubertrick + Segnung 1 AP
			}
			return "/held/page/karmales", ""
		}
	}
	return "", ""
}

func runActionParams(c web.C, w http.ResponseWriter, r *http.Request, action string, params []string) (string, string) {
	switch action {
	case "increment":
		{
			return addToValue(c, w, r, params, 1)
		}
	case "decrement":
		{
			return addToValue(c, w, r, params, -1)
		}
	case "add":
		{
			return addItem(c, w, r, params)
		}
	case "remove":
		{
			return removeItem(c, w, r, params)
		}
	}
	return "", ""
}

func runActionAndRedirect(c web.C, w http.ResponseWriter, r *http.Request) {
	action := c.URLParams["action"]
	params := strings.Split(c.URLParams["*"], "/")
	redirectToURI := ""
	notificationMsg := ""
	if len(params) > 1 {
		redirectToURI, notificationMsg = runActionParams(c, w, r, action, params[1:])
	}
	if PageData.Validator != nil {
		_, PageData.ValidatorMsg = PageData.Validator.Validate()
	}
	calculateAvailable()

	redirInfo := redirectToStruct{RedirectTo: redirectToURI, NotificationMsg: notificationMsg, Magie: PageData.Held.IsMagisch(), Karmal: PageData.Held.IsKarmal(), ValidatorMessages: PageData.ValidatorMsg}

	js, err := json.Marshal(redirInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func isValid(c web.C, w http.ResponseWriter, r *http.Request) {
}

func startPage(c web.C, w http.ResponseWriter, r *http.Request) {
	initPageData()
	renderTemplate(w, "index", &PageData)
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
			eigenschaft.SetMin(eigenschaft.Min() + v.Mod)
			eigenschaft.SetMax(eigenschaft.Max() + v.Mod)
			PageData.Held.Eigenschaften.Set(result, eigenschaft.Wert+v.Mod)
		}
	}

	calculateAvailable()
	renderTemplate(w, "held", &PageData)
}

func runComplexActionAndRedirect(c web.C, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	redirectToURI := ""
	if r.FormValue("type") == "createHeld" {
		newHeld(r)
		// show modification page only if we need it.
		if len(PageData.Held.Spezies.EigenschaftsModifikationen) == 0 {
			redirectToURI = "/held/page/allgemeines"
		} else {
			redirectToURI = "/held/page/modEigenschaften"
		}

	} else if r.FormValue("type") == "modEigenschaften" {
		doModEigenschaften(r)
		redirectToURI = "/held/page/allgemeines"
	}

	calculateAvailable()
	if len(redirectToURI) == 0 {
		return
	}

	redirInfo := redirectToStruct{RedirectTo: redirectToURI, NotificationMsg: "", Magie: PageData.Held.IsMagisch(), Karmal: PageData.Held.IsKarmal(), ValidatorMessages: PageData.ValidatorMsg}

	js, err := json.Marshal(redirInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func newHeld(r *http.Request) {
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
}

func doModEigenschaften(r *http.Request) {
	for i, v := range PageData.Held.Spezies.EigenschaftsModifikationen {
		result := r.FormValue(strconv.Itoa(i))
		eigenschaft := PageData.Held.Eigenschaften.Get(result)
		if eigenschaft != nil {
			eigenschaft.SetMin(eigenschaft.Min() + v.Mod)
			eigenschaft.SetMax(eigenschaft.Max() + v.Mod)
			PageData.Held.Eigenschaften.Set(result, eigenschaft.Wert+v.Mod)
		}
	}
}

// sub pages
func pageNew(c web.C, w http.ResponseWriter, r *http.Request) {
	initPageData()
	if PageData.Held != nil {
		// we have to do stm different if we already have one
		return
	}
	renderTemplate(w, "partials/new", &PageData)
}

func pageModEigenschaften(c web.C, w http.ResponseWriter, r *http.Request) {
	if PageData.Held == nil {
		return // empty page if no held...
	}
	t, _ := template.ParseFiles("template/partials/modEigenschaften.tpl")

	eigenMod := make([]EigenschaftenModSet, len(PageData.Held.Spezies.EigenschaftsModifikationen))
	for i, v := range PageData.Held.Spezies.EigenschaftsModifikationen {
		eigenMod[i].Label = i
		eigenMod[i].Modifikation = v

	}
	t.Execute(w, &eigenMod)
}

func pageAllgemeines(c web.C, w http.ResponseWriter, r *http.Request) {
	if PageData.Held == nil {
		return // empty page if no held...
	}
	renderTemplate(w, "partials/allgemeines", &PageData)
}

func pageKampftechniken(c web.C, w http.ResponseWriter, r *http.Request) {
	if PageData.Held == nil {
		return // empty page if no held...
	}
	renderTemplate(w, "partials/kampftechniken", &PageData)
}

func pageTalente(c web.C, w http.ResponseWriter, r *http.Request) {
	if PageData.Held == nil {
		return // empty page if no held...
	}
	renderTemplate(w, "partials/talente", &PageData)
}

func pageLiturgien(c web.C, w http.ResponseWriter, r *http.Request) {
	if PageData.Held == nil {
		return // empty page if no held...
	}
	renderTemplate(w, "partials/karmales", &PageData)
}

func pageZauber(c web.C, w http.ResponseWriter, r *http.Request) {
	if PageData.Held == nil {
		return // empty page if no held...
	}
	renderTemplate(w, "partials/magie", &PageData)
}

func pageFooter(c web.C, w http.ResponseWriter, r *http.Request) {
	ap_info := apData{AP: 0, AP_spent: 0, AP_total: 0}

	if PageData.Held != nil {
		ap_info.AP = PageData.Held.AP
		ap_info.AP_spent = PageData.Held.AP_spent
		ap_info.AP_total = PageData.Held.AP + PageData.Held.AP_spent

	}
	t, _ := template.ParseFiles("template/partials/footer.tpl")
	if PageData.Held.AP < 0 {
		t, _ = template.ParseFiles("template/partials/footer_negativ.tpl")
	}
	t.Execute(w, &ap_info)
}

// json accessors / REST

func getAP(c web.C, w http.ResponseWriter, r *http.Request) {
	ap_info := apData{AP: 0, AP_spent: 0, AP_total: 0}

	if PageData.Held != nil {
		ap_info.AP = PageData.Held.AP
		ap_info.AP_spent = PageData.Held.AP_spent
		ap_info.AP_total = PageData.Held.AP + PageData.Held.AP_spent
	}

	js, err := json.Marshal(ap_info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func initRoutes() {

	// Setup static files
	static := web.New()
	static.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.Handle("/static/", static)

	// prepare routes, get/post stuff etc
	goji.Get("/", startPage)
	goji.Post("/held/action/:action/*", runActionAndRedirect)
	goji.Post("/held/complexaction", runComplexActionAndRedirect)
	goji.Get("/held/isValid", isValid)
	// partial html stuff - sub-pages
	goji.Get("/held/page/new", pageNew)
	goji.Get("/held/page/modEigenschaften", pageModEigenschaften)
	goji.Get("/held/page/allgemeines", pageAllgemeines)
	goji.Get("/held/page/kampftechniken", pageKampftechniken)
	goji.Get("/held/page/talente", pageTalente)
	goji.Get("/held/page/footer", pageFooter)
	goji.Get("/held/page/karmales", pageLiturgien)
	goji.Get("/held/page/magie", pageZauber)

	// json-accessors/ partial rest-API?
	goji.Get("/held/data/ap", getAP)
}

func Serve(done chan bool) {
	initRoutes()
	goji.Serve()
	done <- true
}
