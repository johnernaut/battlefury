package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnernaut/battlefury/models"
	"github.com/johnernaut/battlefury/db"
    _ "encoding/json"
	"net/http"
)

func BuildRoutes() (*mux.Router, error) {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/", indexHandler)

	return r, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    user := &models.User{}
    people := db.Find(user)

    fmt.Fprint(w, people)
}
