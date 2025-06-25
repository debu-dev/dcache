package utils

import (
	"log"
	"regexp"
)

// IsValidKey checks if the provided key is valid according to specific criteria.
func IsValidKey(key string) bool {
	// Example criteria: key must be alphanumeric and between 1 and 64 characters
	re := regexp.MustCompile(`^[a-zA-Z0-9]{1,64}$`)
	return re.MatchString(key)
}

// LogError logs the provided error message.
func LogError(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}