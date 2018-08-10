package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	const authString = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAO0DRZDndIGc7L8OEXDnzMM4QNVvgjeHzQ3Wwv5NX3sRsxs4Yhx4TAGdtr1WegLzvSiaIIUij20VxgfbHWNMFk="
	var hostKey ssh.PublicKey
	hostKey, _, _, _, _ = ssh.ParseAuthorizedKey([]byte(authString))
	config := &ssh.ClientConfig{
		User: "c2user",
		Auth: []ssh.AuthMethod{
			ssh.Password("c2password"),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	client, error := ssh.Dial("tcp", "172.16.65.137:22", config)
	if error != nil {
		log.Fatal("Failed to connect: ", error)
	}

	session, error := client.NewSession()
	if error != nil {
		log.Fatal("Failed to start session: ", error)
	}

	var b bytes.Buffer
	session.Stdout = &b
	if error := session.Run("/usr/bin/whoami"); error != nil {
		log.Fatal("Failed to run: ", error)
	}
	fmt.Println(b.String())
	session.Close()
}
