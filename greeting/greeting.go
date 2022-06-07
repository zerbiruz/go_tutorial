package greeting

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an erron with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a string that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
