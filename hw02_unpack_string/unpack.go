package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"
	//"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var str strings.Builder
	for i, r := range s {
		if i == 0 && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}

		if (r > 64 && r < 91) || (r > 96 && r < 123) || r == 47 {
			if i == len(s)-1 {
				str.WriteString(string(r))
				return str.String(), nil
			}
			nextRune := rune(s[i+1])
			fmt.Printf("This rune is %s. Next rune is %s \n", string(r), string(nextRune))
			if unicode.IsDigit(nextRune) && unicode.IsDigit(r) {
				return "", ErrInvalidString
			}

			if unicode.IsDigit(nextRune) {
				str.WriteString(strings.Repeat(string(r), int(nextRune-'0')))
			} else {
				str.WriteString(string(r))
			}

			fmt.Println(r)
			fmt.Printf("Index is %d , value is  %s  \n", i, string(r))
		}
	}
	return str.String(), nil
}

// 1. loop string
// 2. check if starts from digit
// 3.

func main() {
	//unicode.IsDigit
	//strings.Builder
	//strings.Repeat
	//strconv.Atoi
	testStr := "a0a0a10bb"

	res, err := Unpack(testStr)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("result", res)
}
