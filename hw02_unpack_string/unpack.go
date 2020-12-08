package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	if isFirstLatterDigit(s) || isConsequentNumbers(s) {
		return "", ErrInvalidString
	}

	builder := strings.Builder{}
	slash := "\\"

	for i, r := range s {
		symbol := string(r)

		if unicode.IsDigit(r) || symbol == slash {
			continue
		}

		if i > 1 && string(s[i-1]) == slash {
			symbol = slash + symbol
		}

		if i < len(s)-1 {
			if count, err := strconv.Atoi(string(s[i+1])); err == nil {
				builder.WriteString(strings.Repeat(symbol, count))

				continue
			}
		}

		builder.WriteString(symbol)
	}

	return builder.String(), nil
}

func isConsequentNumbers(s string) bool {
	previousDigit := false

	for _, r := range s {
		isDigit := unicode.IsDigit(r)

		if isDigit && previousDigit {
			return true
		}

		previousDigit = isDigit
	}

	return false
}

func isFirstLatterDigit(s string) bool {
	return unicode.IsDigit(rune(s[0]))
}
