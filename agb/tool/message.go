package tool

import (
	"fmt"
	"os"
)

// Mnote handles note messages
func Mnote(msg string) {
	fmt.Printf("[ * ] %s", msg)
}

// Mcancel handles cancel messages
func Mcancel(msg string) {
	fmt.Printf("[ - ] %s", msg)
}

// Mdone handle done messages
func Mdone() {
	fmt.Println("Done")
}

// Merror handles generic error messages
func Merror(msg string) {
	fmt.Printf("[ ! ] %s", msg)
	os.Exit(1)
}
