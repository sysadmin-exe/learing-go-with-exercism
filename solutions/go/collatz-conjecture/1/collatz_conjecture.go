package collatzconjecture

import "errors"
import "fmt"


func CollatzConjecture(n int) (int, error) {
	steps := 0
    fmt.Sprintf("input %d", n)
    for n != 1 {
        if n < 0 || n == 0 {
            return 0, errors.New("n is less than or equal to 0")
        }
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = (n * 3) + 1
        }
        steps ++
    }
    return steps, nil
}
