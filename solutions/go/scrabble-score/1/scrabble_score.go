package scrabble

import "fmt"
import "strings"

func Score(word string) int {
	fmt.Println("input string is: ", word)

    count := 0

    s := strings.ToUpper(word)

   for _, char := range s{
		switch char {
            case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
            	count ++
            case 'D', 'G':
            	count = count + 2
            case 'B', 'C', 'M', 'P':
            	count = count + 3
            case 'F', 'H', 'V', 'W', 'Y':
            	count = count + 4 
            case 'K': 
            	count = count + 5
            case 'J', 'X':
            	count = count + 8
            case 'Q', 'Z':
            	count = count + 10
            default:
            	count = count + 0
        }
    }

    return count
}
