package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Name is a combination of first (fname) and last (lname)
type Name struct {
	fname string
	lname string
}

func getInput() string {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}

	return str
}

func getFileAsString(fileName string) string {
	fileContents, e := ioutil.ReadFile(fileName)

	if e != nil {
		fmt.Println(e)
	}

	return string(fileContents)
}

func main() {
	allNames := make([]Name, 0, 1)
	fmt.Println("What file would you like to scan?")
	fileName := getInput()
	fileContents := getFileAsString(fileName)
	names := strings.Split(fileContents, "\n")

	for _, name := range names {
		firstAndLast := strings.Split(name, " ")
		nameObject := Name{
			fname: firstAndLast[0],
			lname: firstAndLast[1],
		}

		allNames = append(allNames, nameObject)
	}

	for _, name := range allNames {
		fmt.Printf("%s %s", name.fname, name.lname)
		fmt.Print("\n")
	}
}
