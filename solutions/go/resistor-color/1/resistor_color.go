package resistorcolor

import "fmt"
// Colors returns the list of all colors.
func Colors() []string {
	colors := []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
	return colors
}

// ColorCode returns the resistance value of the given color.
// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for i, j := range Colors() {
		if j == color {
			fmt.Println(i, j)
			return i
		}
	}
	return 0
}