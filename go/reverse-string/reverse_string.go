package reverse

func Reverse(input string) string {
	var s string
	for _, v := range input {
		s = string(v) + s
	}

	return s
}
