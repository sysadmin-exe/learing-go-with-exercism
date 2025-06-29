package pangram

import (
    "fmt";
    "strings";
)



func IsPangram(input string) bool {
    // create comparable map
    check := map[rune]bool{
        'a':true,
        'b':true,
        'c':true,
        'd':true,
        'e':true,
        'f':true,
        'g':true,
        'h':true,
        'i':true,
        'k':true,
        'j':true,
        'l':true,
        'm':true,
        'n':true,
        'o':true,
        'p':true,
        'q':true,
        'r':true,
        's':true,
        't':true,
        'u':true,
        'v':true,
        'w':true,
        'x':true,
        'y':true,
        'z':true,
    }
    
	fmt.Println("Input Word is:", input)
    // create a map of the values in the input
	seen := make(map[rune]bool)
	for _, r := range strings.ToLower(input) {
		if r < 'a' || r > 'z' {
			continue // Ignore spaces and hyphens
		}
		seen[r] = true
	}
	fmt.Println("Seen map:", seen)
    // compare seen map with check map
	if len(seen) != len(check) {
        fmt.Println("Length not the same")
		return false
	}
	for i, valseen := range seen {
		if valcheck, ok := check[i]; !ok || valseen != valcheck {
            fmt.Println("values not the same")
			return false
		}
	}
	return true
}
