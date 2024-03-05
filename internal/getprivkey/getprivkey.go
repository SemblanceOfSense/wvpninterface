package getprivkey

import (
	"os"
)

type PrivateKeyRequestStruct struct {
    privatekey string
}

func PrivateKeyRequest() ([]byte, error) {
    body, err := os.ReadFile("/etc/wireguard/wg0.conf")
    if err != nil {
        return make([]byte, 0), err
    }

    return body, nil
}
