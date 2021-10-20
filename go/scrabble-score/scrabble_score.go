package scrabble

import (
	"sort"
	"strings"
)

var scoreMap = map[string]int{
	"A": 1,
	"E": 1,
	"I": 1,
	"O": 1,
	"U": 1,
	"L": 1,
	"N": 1,
	"R": 1,
	"S": 1,
	"T": 1,
	"D": 2,
	"G": 2,
	"B": 3,
	"C": 3,
	"M": 3,
	"P": 3,
	"F": 4,
	"H": 4,
	"V": 4,
	"W": 4,
	"Y": 4,
	"K": 5,
	"J": 8,
	"X": 8,
	"Q": 10,
	"Z": 10,
}

// Score returns the score of a word for a game of Scrabble.
func Score(word string) int {

	wordSlice := strings.Split(strings.ToUpper(word), "")
	letterCounts := countDuplicates(wordSlice)

	wordScore := 0
	for k, v := range letterCounts {
		wordScore += scoreMap[k] * v
	}

	return wordScore
}

// countDuplicates takes a slice of strings and sorts them
// once sorted the function counts any duplicated values returning a
// map with the value and number of occurences of that character.
func countDuplicates(s []string) map[string]int {
	sort.Strings(s)
	duplicateCount := make(map[string]int)

	for _, v := range s {
		duplicateCount[v] += 1
	}

	return duplicateCount
}
