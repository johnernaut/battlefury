package main

import (
    bf "github.com/johnernaut/battlefury"
    "log"
    "net/http"
)

func main() {
    r, err := bf.BuildRoutes()

    if err != nil {
        log.Panic(err)
    }

    http.Handle("/", r)
    log.Println("Listening on port 1004")
    log.Fatal(http.ListenAndServe(":1004", http.DefaultServeMux)
}
