package main

import (
	"fmt"

	"irh/greet"
)

func main() {
	message := greet.Hello("Nacho")
	fmt.Println(message)
}
