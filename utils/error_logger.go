package utils

import (
	"log"
)

// LogError will log error clearly
func LogError(message string, err error) {
	log.Printf("\n\n%v\nError: %v\n\n", message, err)
}
