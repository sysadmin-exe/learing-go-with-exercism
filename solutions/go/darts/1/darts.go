package darts

import (
    "fmt";
    "math"
)


func Score(x, y float64) int {
	fmt.Println("Input coordinates:", x, y)

    // find the radius using pythagoras
    r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
    fmt.Println("Input radius:", r)

	// outer ?
    if  r <= 10 && r > 5 {
        fmt.Println("1")
        return 1
    }
    // middle ?
    if r <= 5 && r > 1 {
        fmt.Println("5")
        return 5
    }
    if r <= 1 && r >= 0 {
        fmt.Println("10")
        return 10
    } 
	return 0
}
