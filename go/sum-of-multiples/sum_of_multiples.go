package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var s int

	var mult = make(map[int]bool)

	for _, v := range divisors {
		if v == 0 {
			continue
		}
		for i := v; i < limit; i++ {
			if i%v == 0 {
				if _, ok := mult[i]; !ok {
					mult[i] = true
					s += i
				}
			}
		}
	}

	return s
}
