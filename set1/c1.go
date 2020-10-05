package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

// HexToBase64 converts a hex string to a base64 string
func HexToBase64(hs string) (string, error) {
	v, err := hex.DecodeString(hs)
	if err != nil {
		return "", err
	}
	log.Printf("Decoded input: %s", v)
	return base64.StdEncoding.EncodeToString(v), nil
}
