package luhn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")

	if len(s) <= 1 {
		return false
	}

	for _, v := range s {
		if !unicode.IsNumber(v) {
			return false
		}
	}

	for i := len(s) - 2; i >= 0; i -= 2 {
		n, err := strconv.Atoi(string(s[i]))

		if err != nil {
			panic(err)
		}
		product := n * 2
		if product > 9 {
			product -= 9
		}

		s = s[:i] + fmt.Sprint(product) + s[i+1:]
	}

	sum := 0
	for _, v := range s {
		n, err := strconv.Atoi(string(v))

		if err != nil {
			panic(err)
		}

		sum += n
	}

	return sum%10 == 0
}
