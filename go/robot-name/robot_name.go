package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var validNames = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	for a := 'A'; a <= 'Z'; a++ {
		for b := 'A'; b <= 'Z'; b++ {
			for c := 0; c <= 999; c++ {
				validNames[fmt.Sprintf("%c%c%.3d", a, b, c)] = false
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

	for k := range validNames {
		r.name = k
		delete(validNames, k)
		break
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
