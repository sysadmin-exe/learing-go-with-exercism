package reverse

import "fmt"

func Reverse(input string) string {
	fmt.Println("Input:", input)

	runes := []rune(input) // convert to runes for Unicode safety
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
