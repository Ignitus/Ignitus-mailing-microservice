package utils

import (
	"log"
)

// LogMessage is used to log without looking messy.
func LogMessage(msg interface{}) {
	log.Printf("\n\n%v\n\n", msg)
}
