package greeting

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
	// Return a string that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
