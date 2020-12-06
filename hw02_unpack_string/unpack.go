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

	if isFirstLatterDigit(s) || isOnlyDigit(s) || isExistNumber(s) {
		return "", ErrInvalidString
	}

	builder := strings.Builder{}

	for i, r := range s {
		if unicode.IsDigit(r) {
			continue
		}

		if i < len(s)-1 {
			if count, err := strconv.Atoi(string(s[i+1])); err == nil {
				builder.WriteString(strings.Repeat(string(r), count))

				continue
			}
		}

		builder.WriteString(string(r))
	}

	return builder.String(), nil
}

func isExistNumber(s string) bool {
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

func isOnlyDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func isFirstLatterDigit(s string) bool {
	return unicode.IsDigit(rune(s[0]))
}
