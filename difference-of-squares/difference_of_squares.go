package diffsquares

func SquareOfSum(n int) (square int) {
	for i := n; i >= 1; i-- {
		square += i
	}

	square *= square

	return
}

func SumOfSquares(n int) (square int) {
	for i := n; i >= 1; i-- {
		square += i * i
	}

	return
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
