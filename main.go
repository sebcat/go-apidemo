package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	listen = flag.String("listen", ":8080", "listening directive")
	htdocs = flag.String("htdocs", "htdocs", "path to htdocs")
)

// Application for storage and retrieval of strings
type DemoAPI struct {
	store []string
	m     sync.RWMutex
}

// Create a new instance of DemoAPI
//
// Wrapping the allocation of DemoAPI inside a function allows
// us to initialize fields in the future, if we need to
func NewDemoAPI() *DemoAPI {
	return &DemoAPI{}
}

// Save a string in the app
func (app *DemoAPI) save(str string) {
	app.m.Lock()
	app.store = append(app.store, str)
	app.m.Unlock()
}

// Write all stored strings to a writer
func (app *DemoAPI) write(w io.Writer) error {
	enc := json.NewEncoder(w)
	app.m.RLock()
	defer app.m.RUnlock()
	return enc.Encode(app.store)
}

// HTTP handler implementation
func (app *DemoAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("str")
	if r.Method == "POST" && len(str) > 0 {
		app.save(str)
		http.Redirect(w, r, "../", http.StatusFound)
	} else if r.Method == "GET" {
		if err := app.write(w); err != nil {
			log.Println(r.RemoteAddr, r.URL, err)
		}
	}
}

func main() {
	flag.Parse()
	h := NewDemoAPI()
	http.Handle("/api/1", h)
	http.Handle("/", http.FileServer(http.Dir(*htdocs)))
	log.Println("Starting application at", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
