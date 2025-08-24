package diffsquares
import "math"

func SquareOfSum(n int) int {
	var sum int = (n * (n+1))/2
    return sum * sum
}

func SumOfSquares(n int) int {
	var sum int = (n*(n+1)*(2*n+1))/6
    return sum
}

func Difference(n int) int {
	square_of_sum := SquareOfSum(n)
    sum_of_square:= SumOfSquares(n)
    return int(math.Abs(float64(square_of_sum - sum_of_square)))
}
