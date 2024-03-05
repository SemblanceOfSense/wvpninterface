package getprivkey

import (
	"encoding/json"
	"os"
	"strings"
)

type PrivateKeyRequestStruct struct {
    privatekey string
}

func PrivateKeyRequest() ([]byte, error) {
    requstStruct := &PrivateKeyRequestStruct{}

    body, err := os.ReadFile("/etc/wireguard/wg0.conf")
    if err != nil {
        return make([]byte, 0), err
    }

    p := substr(string(body), indexAt(string(body), "y = ", 0) + 4, (indexAt(string(body), "g=", 0) + 2) - (indexAt(string(body), "y = ", 0) + 4))

    requstStruct.privatekey = p

    j, err := json.Marshal(p)
    if err != nil { return make([]byte, 0), err }

    return j, nil
}

func indexAt(s, sep string, n int) int {
    idx := strings.Index(s[n:], sep)
    if idx > -1 {
        idx += n
    }
    return idx
}

func substr(input string, start int, length int) string {
    asRunes := []rune(input)

    if start >= len(asRunes) {
        return ""
    }

    if start+length > len(asRunes) {
        length = len(asRunes) - start
    }

    return string(asRunes[start : start+length])
}
