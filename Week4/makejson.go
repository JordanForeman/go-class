package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func getInput() string {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}

	return str
}

func main() {
	type Person struct {
		Name    string
		Address string
	}

	fmt.Println("What is your name?")
	var name string = getInput()
	fmt.Println("What is your address?")
	var address string = getInput()

	user := Person{
		Name:    name,
		Address: address,
	}

	userJSON, e := json.Marshal(user)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(userJSON))
}
