package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func BuildRoutes() (*mux.Router, error) {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/", indexHandler)

	return r, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}
