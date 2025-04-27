package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func ContentHash(content []byte) (string, error) {
	var hasher = sha256.New()
	_, err := hasher.Write(content)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString((hasher.Sum(nil))), nil
}

func ContentHashV2(content io.Reader) (string, error) {

	var hasher = sha256.New()
	if _, err := io.Copy(hasher, content); err != nil {
		return "", err
	}

	return hex.EncodeToString((hasher.Sum(nil))), nil
}
