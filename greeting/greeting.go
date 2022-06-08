package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an erron with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a string that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met.",
	}

	// Return a randomly selected message format by specifying
	// a random random index for for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
