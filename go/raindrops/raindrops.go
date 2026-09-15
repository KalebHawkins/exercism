package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var buf strings.Builder

	if number%3 == 0 {
		buf.WriteString("Pling")
	}
	if number%5 == 0 {
		buf.WriteString("Plang")
	}
	if number%7 == 0 {
		buf.WriteString("Plong")
	}

	if buf.String() != "" {
		return buf.String()
	}
	return strconv.Itoa(number)
}
