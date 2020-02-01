package main

import (
	"fmt"
	"sort"
	"strconv"
)

func prompt(count int) {
	if count > 1 {
		fmt.Println("Give me another integer (or x to exit)")
	} else {
		fmt.Println("Give me an integer (or x to exit)")
	}
}

func main() {
	homeSlice := make([]int, 0, 3)

	var count int
	for {
		count++

		var currentInput string

		prompt(count)
		fmt.Scan(&currentInput)

		if currentInput == "x" {
			fmt.Println("Thanks for playing!")
			return
		}

		currentInputAsInt, err := strconv.Atoi(currentInput)

		if err != nil {
			fmt.Println("I said an integer! That's not an integer!")
			continue
		}

		homeSlice = append(homeSlice, currentInputAsInt)

		sort.Ints(homeSlice)

		fmt.Println("Current Integers:")
		fmt.Println(homeSlice)
	}
}
