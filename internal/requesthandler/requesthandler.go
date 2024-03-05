package requesthandler

import (
    "io"
    "log"
    "fmt"
    "encoding/json"
)

type AddPeerRequestStruct struct {
    Publickey string
}

func AddPeerRequest(v io.ReadCloser) AddPeerRequestStruct {
    dec := json.NewDecoder(v)

    var request AddPeerRequestStruct
    err := dec.Decode(&request)
    if err != nil {
        fmt.Println("Failed decode")
        log.Fatal(err)
    }

    return request
}
