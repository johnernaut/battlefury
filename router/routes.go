package router

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func BuildRoutes() (*mux.Router, error) {
    r := mux.NewRouter()
    r.HandleFunc("/", indexHandler)

    return r, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}
