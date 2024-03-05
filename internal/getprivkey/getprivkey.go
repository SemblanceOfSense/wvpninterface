package getprivkey

import (
	"os"
    "strings"
)

type PrivateKeyRequestStruct struct {
    privatekey string
}

func PrivateKeyRequest() ([]byte, error) {
    body, err := os.ReadFile("/etc/wireguard/wg0.conf")
    if err != nil {
        return make([]byte, 0), err
    }

    return []byte(substr(string(body), indexAt(string(body), "y = ", 0) + 4, (indexAt(string(body), "g=", 0) + 2) - (indexAt(string(body), "y = ", 0) + 4))), nil
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
