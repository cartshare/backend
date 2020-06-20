package main

import (
	"crypto/rand"
	"encoding/hex"
)

func genToken(lenTarget int) (string, error) {
	// Generate savage 64 byte session token

	tok := []byte{}

	for len(tok) < lenTarget {
		temp := make([]byte, 8)

		n, err := rand.Read(temp)

		if err != nil {
			return "", err
		}

		temp = temp[:n]

		tok = append(tok, temp...)
	}

	tok = tok[:lenTarget-1]

	tokStr := hex.EncodeToString(tok)

	return tokStr, nil
}
