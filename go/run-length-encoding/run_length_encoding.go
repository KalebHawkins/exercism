package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	var encoded string

	for len(input) > 0 {
		r := input[0]
		originalLen := len(input)
		input = strings.TrimLeft(input, string(r))
		if count := originalLen - len(input); count > 1 {
			encoded += fmt.Sprintf("%d", count)
		}
		encoded += string(r)
	}

	return encoded
}

func RunLengthDecode(input string) string {
	var decoded string
	for len(input) > 0 {
		i := strings.IndexFunc(input, func(r rune) bool {
			return !unicode.IsDigit(r)
		})

		n := 1
		if i != 0 {
			n, _ = strconv.Atoi(input[:i])
		}

		decoded += strings.Repeat(string(input[i]), n)
		input = input[i+1:]
	}

	return decoded
}
