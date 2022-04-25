package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var results []string

	for _, s := range candidates {
		// If the length of the strings are the same we can safely assume they are
		// not anagrams. The loop is continued.
		if len(s) != len(subject) {
			continue
		}

		// If the two words are the same they are not anagrams they are jus the
		// same word. The loop is continued.
		if strings.ToLower(s) == strings.ToLower(subject) {
			continue
		}

		// If we turn the canidate string `s` into a byteSlice and sort it
		// and perform the same function with the subject, we can compare the two resulting
		// slices as their string counter part to see if they are the same.
		byteSlice := []byte(strings.ToLower(s))
		sort.Slice(byteSlice, func(i, j int) bool {
			return byteSlice[i] < byteSlice[j]
		})

		subjectSlice := []byte(strings.ToLower(subject))

		sort.Slice(subjectSlice, func(i, j int) bool {
			return subjectSlice[i] < subjectSlice[j]
		})

		// If the resulting slices are the same then we have an anagram.
		// The canidate string is added to the results and the loop continues.
		if string(byteSlice) == string(subjectSlice) {
			results = append(results, s)
		}
	}

	// Once the canidates have been exhausted we return the resulting anagrams.
	return results
}
