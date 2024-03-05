package getprivkey

import (
	"encoding/json"
	"os"
)

type PrivateKeyRequestStruct struct {
    privatekey string
}

func PrivateKeyRequest() ([]byte, error) {
    requstStruct := &PrivateKeyRequestStruct{}

    body, err := os.ReadFile("/etc/wireguard/public.key")
    if err != nil {
        return make([]byte, 0), err
    }

    p := string(body)
    requstStruct.privatekey = p

    j, err := json.Marshal(p)
    if err != nil { return make([]byte, 0), err }

    return j, nil
}
