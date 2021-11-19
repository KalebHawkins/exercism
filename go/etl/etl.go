package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	nm := make(map[string]int)

	for k, v := range in {
		for _, sv := range v {
			nm[strings.ToLower(sv)] = k
		}
	}

	return nm
}
