package isogram


import (
    "fmt";
    "strings";
)

func IsIsogram(word string) bool {
	fmt.Println("Input Word is:", word)

	seen := make(map[rune]bool)
	for _, r := range strings.ToLower(word) {
		if r == ' ' || r == '-' {
			continue // Ignore spaces and hyphens
		}
		if seen[r] {
			return false // Found a repeating letter
		}
		seen[r] = true
	}
	return true
}
