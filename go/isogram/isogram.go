package isogram

import (
	"sort"
	"strings"
)

func IsIsogram(phrase string) bool {
	if phrase == "" {
		return true
	}

	phrase = strings.ToUpper(phrase)
	phrase = strings.ReplaceAll(phrase, "-", "")
	phrase = strings.ReplaceAll(phrase, " ", "")

	phraseSlice := strings.Split(phrase, "")
	sort.Strings(phraseSlice)

	for i := 0; i < len(phraseSlice)-1; i++ {
		if phraseSlice[i] == phraseSlice[i+1] {
			return false
		}
	}

	return true
}
