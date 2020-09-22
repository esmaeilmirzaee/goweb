package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"webvideos/060/hash"
)

func main() {
	// Checking HMAC
	toHash := []byte("This is my another secret key")
	h := hmac.New(sha256.New, []byte("This is my secret key"))
	h.Write(toHash)
	b := h.Sum(nil)
	fmt.Println(base64.URLEncoding.EncodeToString(b))

	a := hash.NewHMAC("This is my secret key")
	fmt.Println(a.Hash("This is my another secret key"))
}
