package main

import (
	"fmt"

	"example.com/greeting"
)

func main() {
	message := greeting.Hello("Theera")
	fmt.Println(message)
}
