package getpubkey

import (
	"encoding/json"
	"os"
)

type PublicKeyRequestStruct struct {
    publickey string
}

func PublicKeyRequest() ([]byte, error) {
    requstStruct := &PublicKeyRequestStruct{}

    body, err := os.ReadFile("/etc/wireguard/public.key")
    if err != nil {
        return make([]byte, 0), err
    }

    p := string(body)
    inputFmt:=p[:len(p)-1]
    requstStruct.publickey = inputFmt

    j, err := json.Marshal(p)
    if err != nil { return make([]byte, 0), err }

    return j, nil
}
