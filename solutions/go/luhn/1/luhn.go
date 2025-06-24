package luhn

import (
    "fmt";
    "strings";
    )

func Valid(id string) bool {
	fmt.Println("Input:", id)

	// Remove spaces
	cleaned := strings.ReplaceAll(id, " ", "")

    result := []int{}

	// Convert string to slice of digits
	for _, r := range cleaned {
		result = append(result, int(r-'0'))
	}

    fmt.Println("CleanedResult:", result)
    
    // check if input is valid
    if len(result) == 1 {
        if result[0] == 0 {
            return false
        }
    }

    // check for ascii characters
	for _, r := range result {
		if r >= 10 { 
			return false
		}
	}
    
	// Add 2 to every second digit from the right
	for i := len(result) - 2; i >= 0; i -= 2 {
		result[i] += result[i]
        if result[i] > 9 {
            result[i] -= 9
        } else {
            result[i] = result[i]
        }
	}

	fmt.Println("DoubledResult:", result)


    sum := 0
	for _, j := range result {
		sum += j
	}
	fmt.Println("Sum:", sum)

    if sum % 10 == 0 {
        fmt.Println("Valid?: yes")
        return true
    } else {
        fmt.Println("Valid?: no")
        return false
    }
}

