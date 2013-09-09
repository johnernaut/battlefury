package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnernaut/battlefury/models"
	"net/http"
)

func BuildRoutes() (*mux.Router, error) {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")

	r.HandleFunc("/", indexHandler)
    r.HandleFunc("/users/{id:[0-9]+}", userHandler)

	return r, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    user := &models.User{}
    people := user.All()

    fmt.Fprint(w, string(people))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    user := &models.User{}
    people := user.Find(id)

    fmt.Fprint(w, string(people))
}
