package twelve

import "fmt"

var theDays = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var theGifts = []string{"a Partridge in a Pear Tree.", "two Turtle Doves, ", "three French Hens, ", "four Calling Birds, ", "five Gold Rings, ", "six Geese-a-Laying, ", "seven Swans-a-Swimming, ", "eight Maids-a-Milking, ", "nine Ladies Dancing, ", "ten Lords-a-Leaping, ", "eleven Pipers Piping, ", "twelve Drummers Drumming, "}
var firstPart = "On the %s day of Christmas my true love gave to me: "

func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i)
		if i != 12 {
			song += "\n"
		}
	}

	return song
}

func Verse(i int) string {
	var verse string
	verse = fmt.Sprintf(firstPart, theDays[i-1])

	for day := i - 1; day >= 0; day-- {
		if i > 1 && day == 0 {
			verse += "and "
		}
		verse += theGifts[day]
	}

	return verse
}
