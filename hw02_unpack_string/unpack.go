package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var str strings.Builder
	for i, r := range s {
		if i == 0 && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}

		if i == len(s)-1 {
			if !unicode.IsDigit(r) {
				str.WriteRune(r)
			}
			return str.String(), nil
		}

		nextRune := rune(s[i+1])
		if unicode.IsDigit(nextRune) && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(r) {
			continue
		}

		if unicode.IsDigit(nextRune) {
			str.WriteString(strings.Repeat(string(r), int(nextRune-'0')))
		} else {
			str.WriteRune(r)
		}
	}
	return str.String(), nil
}
