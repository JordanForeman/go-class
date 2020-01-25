package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const I string = "I"
const A string = "A"
const N string = "N"
const FOUND string = "Found!"
const NOT_FOUND string = "Not Found!"

func startsWith(src string, char string) bool {
	return strings.HasPrefix(src, char)
}

func endsWith(src string, char string) bool {
	return strings.HasSuffix(src, char)
}

func contains(src string, char string) bool {
	return strings.Contains(src, char)
}

/*
The following method uses an alternative to fmt.scan to support
scanning a single string from user input while maintaining spaces.

see: https://stackoverflow.com/a/34647156/782884

It is unclear whether or not this is in the spirit of the Coursera course
requirements, and is discussed (without resolution) in the assignment's
discussion tab.

see: https://www.coursera.org/learn/golang-getting-started/peer/f1Cly/module-2-activity-findian-go/discussions/threads/YysytfB3Eeij7g78JfEN-A

*/
func getInput() string {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}

	return str
}

func main() {
	fmt.Println("Help me find Ian")

	str := getInput()
	big := strings.ToUpper(str)

	if startsWith(big, I) && contains(big, A) && endsWith(big, N) {
		fmt.Println(FOUND)
	} else {
		fmt.Println(NOT_FOUND)
	}
}
