package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	m := make(map[string]int)
	re := regexp.MustCompile(`[\w']+`)
	words := re.FindAllString(strings.ToLower(phrase), -1)

	for _, word := range words {
		word = strings.Trim(word, "'")
		m[word]++
	}

	return m
}
