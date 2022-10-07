package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {

	var str strings.Builder

	for _, r := range plain {
		if !unicode.IsLetter(r) {
			str.WriteRune(r)
			continue
		}
		cipheredRune := shiftKey + int(r)

		if cipheredRune > 90 && unicode.IsUpper(r) {
			cipheredRune -= 26
		}

		if cipheredRune > 122 && unicode.IsLower(r) {
			cipheredRune -= 26
		}

		str.WriteRune(rune(cipheredRune))
	}

	return str.String()
}
