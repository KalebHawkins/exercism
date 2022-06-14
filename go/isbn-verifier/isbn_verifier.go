package isbn

import (
	"unicode"
)

func IsValidISBN(isbn string) bool {
	chk := 0
	cnt := 0

	for _, x := range isbn {
		switch {
		case unicode.IsDigit(x):
			chk += (10 - cnt) * int(x-'0')
			cnt++
		case x == 'X' && cnt == 9:
			chk += 10
			cnt++
		}

	}

	return cnt == 10 && chk%11 == 0

}
