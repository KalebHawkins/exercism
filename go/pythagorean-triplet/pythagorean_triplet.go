package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	res := []Triplet{}

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					res = append(res, Triplet{a, b, c})
				}
			}
		}
	}

	return res
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	res := make([]Triplet, 0)

	for _, t := range Range(1, p/2) {
		if t[0]+t[1]+t[2] == p {
			res = append(res, t)
		}
	}

	return res
}
