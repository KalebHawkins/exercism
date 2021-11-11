package proverb

import "fmt"

// Proverb generates a proverb based on For want of a nail: http://en.wikipedia.org/wiki/For_Want_of_a_Nail
func Proverb(rhyme []string) []string {
	formatString := "For want of a %v the %v was lost."
	finalVerse := "And all for the want of a %v."
	verses := make([]string, 0)

	for i := range rhyme {
		if i != len(rhyme)-1 {
			verses = append(verses, fmt.Sprintf(formatString, rhyme[i], rhyme[i+1]))
		} else {
			verses = append(verses, fmt.Sprintf(finalVerse, rhyme[0]))
		}

	}

	return verses
}
