package robotname

import (
	"errors"
	"fmt"
)

type Robot struct {
	name string
}

var validNames = make([]string, 26*26*10*10*10)

func init() {
	var index int = 0
	for a := 'A'; a <= 'Z'; a++ {
		for b := 'A'; b <= 'Z'; b++ {
			for c := 0; c <= 999; c++ {
				validNames[index] = fmt.Sprintf("%c%c%.3d", a, b, c)
				index++
			}
		}
	}
}
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(validNames) == 0 {
		return "", errors.New("namespace exhausted: no more valid names left")
	}
	r.name = validNames[0]
	validNames = validNames[1:]
	return r.name, nil
}
func (r *Robot) Reset() {
	r.name = ""
}
