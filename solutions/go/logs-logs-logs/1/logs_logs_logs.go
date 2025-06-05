package logs

import "unicode/utf8"
// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
    	if char == '❗' {
            return "recommendation"
        } 
        if char == '🔍' {
        	return "search"
    	} 
        if char == '☀' {
        	return "weather"
        } 
        continue
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    runes := []rune(log)
	for index, char := range runes {
        if char == oldRune {
        	runes[index] = newRune
        }
    }
    return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)
    if numberOfRunes <= limit {
        return true
    } else {
        return false
    }
}
