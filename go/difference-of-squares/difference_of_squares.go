package diffsquares

func SquareOfSum(n int) int {
	var value int
	for i := 0; i <= n; i++ {
		value += i
	}

	return value * value
}

func SumOfSquares(n int) int {
	var value int
	for i := 0; i <= n; i++ {
		value += i * i
	}

	return value
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
