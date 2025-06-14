package hamming

import "errors"
import "fmt"

func Distance(a, b string) (int, error) {
    fmt.Println("string a is: ", a)
    fmt.Println("string b is: ", b)

    diff := 0

    runes1 := []rune(a)
    runes2 := []rune(b)

    if len(runes1) != len(runes2) {
        fmt.Println("diff is: ", diff)
        return diff, errors.New("strings not the same")
    }

    maxLen := len(runes1)
    if len(runes2) > maxLen {
        maxLen = len(runes2)
    }
    
    for i := 0; i < maxLen; i++ {
		if runes1[i] != runes2[i] {
            diff++
        }
    }

    fmt.Println("diff is: ", diff)
    return diff, nil
}
