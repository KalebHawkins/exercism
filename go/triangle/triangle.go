package triangle

type Kind string

const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if (a <= 0 && b <= 0 && c <= 0) || (a+b <= c || b+c <= a || a+c <= b) {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else if a != b && a != c && b != c {
		k = Sca
	}

	return k
}
