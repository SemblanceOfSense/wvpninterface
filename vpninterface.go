package main

import (
	"fmt"
	"log"
	"net/http"
	"vpninterface/internal/addpeer"
)

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/":
        fmt.Println("hi")
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

type HelloHandler struct {}

func main() {
    err := addpeer.AddPeer("eeeeee")
    if err != nil { log.Fatal(err) }
    http.Handle("/", &HelloHandler{})

    log.Fatal(http.ListenAndServe(":8080", nil))
}
