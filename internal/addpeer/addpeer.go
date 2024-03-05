package addpeer

import (
    "os"
)

func AddPeer(publickey string) error {
    file, err := os.OpenFile("/etc/wireguard/wg0.conf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil { return err }

    newpeer := "\n[Peer]\nPublicKey = " + publickey + "\n\nAllowedIPs = 0.0.0.0/0\n\nPersistentKeepalive = 25"

    if _, err := file.Write([]byte(newpeer)); err != nil {
        file.Close()
        return err
    }
    if err := file.Close(); err != nil {
        return err
    }

    return nil
}
