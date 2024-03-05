package getpubkey

import (
	"encoding/json"
	"os"
    "strings"
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
    p = strings.TrimSuffix(p, "\n")
    requstStruct.publickey = p

    j, err := json.Marshal(p)
    if err != nil { return make([]byte, 0), err }

    return j, nil
}
