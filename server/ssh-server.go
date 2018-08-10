package main

import (
     "fmt"
     "log"

     "golang.org/x/crypto/ssh"
)

func main() {
     // Create very basic SSH server config
     config := &ssh.ServerConfig{
          PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
               if c.User() == "c2user" && string(pass) == "c2password" {
                    return nil, nil
               }
               return nil, fmt.Errorf("Password rejected, %q", c.User())
          },
     }

     privateKeyString = "-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIBgXU0TxInvi67PusMrMyWolN1z4BFx8AeN/zWXj2hXBbl42pBU/dM
ZED/W9ZqjpsZUKF3vKNZ0hFROoztozx1+mKgBwYFK4EEACOhgYkDgYYABABdmQT3
66DyKuNmRpTZC+R/dMq4e9+2Y1ytoe7ytrg4rL5TxZOcac1wieUIs0wQv6FHeajk
ZTCyLqqsFJ5d8xXYagB5b2nlLJZIKf56TlAmgHgX2AelXkmCmWtnBMSWXTOKyF+1
Uv/Vmfjc4SDQ6OPt0BNWTIP3t70Y64yK4ouUAigruA==
-----END EC PRIVATE KEY-----"
     privateKey, error = ssh.ParsePrivateKey([]byte(privateKeyString))

     config.AddHostKey(privateKey)

     // Begin accepting connections

}
