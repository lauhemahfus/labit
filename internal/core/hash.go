package core

import (
	"crypto/sha256"
	"fmt"
)

func HashContent(content []byte) string {
	hash := sha256.Sum256(content)
	return fmt.Sprintf("%x", hash)
}

func HashString(s string) string {
	return HashContent([]byte(s))
}
