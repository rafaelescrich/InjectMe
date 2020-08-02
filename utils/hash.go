package utils

import (
	"crypto/sha256"
	"fmt"
)

// Hash takes a string and returns it's sha2-256 hash
func Hash(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}
