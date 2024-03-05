package main

import (
	"fmt"
	"log"
	"net/http"
	"vpninterface/internal/addpeer"
	"vpninterface/internal/getprivkey"
	"vpninterface/internal/requesthandler"
)

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/addpeer":
        switch r.Method {
        case http.MethodPost:
            err := addpeer.AddPeer(requesthandler.AddPeerRequest(r.Body).Publickey)
            if err != nil { fmt.Println(err) }
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    case "/getprivkey":
        switch r.Method {
        case http.MethodGet:
            s, err := getprivkey.PrivateKeyRequest()
            if err != nil { fmt.Println(err) }
            w.Write(s)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

type HelloHandler struct {}

func main() {
    http.Handle("/", &HelloHandler{})

    log.Fatal(http.ListenAndServe(":8080", nil))
}
