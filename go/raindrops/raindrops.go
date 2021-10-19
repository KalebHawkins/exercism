package raindrops

import "fmt"

func Convert(n int) string {
	results := ""

	if n%3 == 0 {
		results += "Pling"
	}
	if n%5 == 0 {
		results += "Plang"
	}
	if n%7 == 0 {
		results += "Plong"
	}

	if results == "" {
		results = fmt.Sprintf("%v", n)
	}

	return results
}
