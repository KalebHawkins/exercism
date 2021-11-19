package pangram

import "strings"

var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsPangram(input string) bool {
	input = strings.ToUpper(input)

	for _, r := range alphabet {
		if !strings.ContainsRune(input, r) {
			return false
		}
	}

	return true
}
