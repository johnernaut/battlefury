package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnernaut/battlefury/models"
    "encoding/json"
	"net/http"
)

func BuildRoutes() (*mux.Router, error) {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/", indexHandler)

	return r, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    people := models.Find()
    data, err := json.Marshal(people)

    if err != nil {
        panic(err)
    }

    fmt.Fprint(w, string(data))
}
