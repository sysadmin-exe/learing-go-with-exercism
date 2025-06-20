package diffsquares

import "fmt"

func SquareOfSum(n int) int {
	fmt.Println("Input:", n)

    sqofsu := 0
    for i := 0; i <= n; i++ {
        sqofsu = i + sqofsu
    }
    return sqofsu * sqofsu
}

func SumOfSquares(n int) int {
	fmt.Println("Input:", n)

    suofsq := 0
	for i := 0; i <= n; i++ {
        suofsq = (i * i) + suofsq
    }
    return suofsq
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
