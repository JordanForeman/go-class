package main

import "fmt"

func main() {
	var userInput float32

	fmt.Print("Give me a decimal please:\n")
	fmt.Scan(&userInput)
	fmt.Printf("Truncated: %d", int32(userInput))
}
