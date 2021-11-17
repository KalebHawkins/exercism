package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var rtn Ints

	for _, v := range i {
		if filter(v) {
			rtn = append(rtn, v)
		}
	}

	return rtn
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var rtn Ints

	for _, v := range i {
		if !filter(v) {
			rtn = append(rtn, v)
		}
	}

	return rtn
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var rtn Lists

	for _, v := range l {
		if filter(v) {
			rtn = append(rtn, v)
		}
	}

	return rtn
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var rtn Strings

	for _, v := range s {
		if filter(v) {
			rtn = append(rtn, v)
		}
	}

	return rtn
}
