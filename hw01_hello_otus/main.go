package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	testString := "Hello, OTUS!"
	printReversedString(testString)
}

func printReversedString(str string) {
	fmt.Println(stringutil.Reverse(str))
}
