package rand

import (
	"crypto/rand"
	"encoding/base64"
)

// RememberTokenBytes should be greater than 32.
const RememberTokenBytes = 32

// Byte generates a slice of random bytes.
// If there is err it would return nil and error.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, err
}

// String converts created random number to string,
// or return error if there is any exist.
func String(nByte int) (string, error) {
	b, err := Bytes(nByte)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken return a string or an error.
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}
