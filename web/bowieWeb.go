package bowieweb

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

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

func initRoutes() {
	// prepare routes, get/post stuff etc
	//e.g. goji.Post("/held/reset", resetHero) // initial state (select name/species/culture...)
	goji.Post("/held/reset", resetHero)
	goji.Post("/held/increment/:name", incrementValue)
	goji.Post("/held/decrement/:name", decrementValue)
	goji.Get("/held/canincrement/:name", canIncrementValue) // true if enough AP available
	goji.Get("/held/candecrement/:name", canDecrementValue) // true if not at min value
	goji.Post("/held/set/:value/*", setValue)
	goji.Get("/held/get/:value/*", getValue)
	goji.Get("/held/isValid", isValid)
}

func Serve() {
	//goji.Get(pattern, handler)
	initRoutes()
	goji.Serve()
}

/*func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Serve()"
)*/
