package main

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func main() {
	// This key will need to be updated in the future, since this is a test key
	const authString = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAO0DRZDndIGc7L8OEXDnzMM4QNVvgjeHzQ3Wwv5NX3sRsxs4Yhx4TAGdtr1WegLzvSiaIIUij20VxgfbHWNMFk="
	var hostKey ssh.PublicKey
	hostKey, _, _, _, _ = ssh.ParseAuthorizedKey([]byte(authString))
	// Set up connection config
	config := &ssh.ClientConfig{
		User: "c2user",
		Auth: []ssh.AuthMethod{
			ssh.Password("c2password"),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// Connect to host
	client, error := ssh.Dial("tcp", "172.16.65.137:443", config)
	if error != nil {
		return
	}

	// Open shell
	session, error := client.NewSession()
	if error != nil {
		return
	}

	// Run whoami, print results, exit
	var b bytes.Buffer
	session.Stdout = &b
	if error := session.Run("/usr/bin/whoami"); error != nil {
		return
	}
	fmt.Println(b.String())
	session.Close()
}
