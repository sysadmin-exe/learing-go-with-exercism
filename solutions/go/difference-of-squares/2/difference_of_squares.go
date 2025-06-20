package diffsquares

import "fmt"

func SquareOfSum(n int) int {
	fmt.Println("Input:", n)

    sqofsu := 0
    for i := 0; i <= n; i++ {
        sqofsu += i
    }
    return sqofsu * sqofsu
}

func SumOfSquares(n int) int {
	fmt.Println("Input:", n)

    suofsq := 0
	for i := 0; i <= n; i++ {
        suofsq += (i * i)
    }
    return suofsq
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
