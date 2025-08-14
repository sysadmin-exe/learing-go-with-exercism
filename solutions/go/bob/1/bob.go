// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
    "fmt";
    s "strings"
    u "unicode"
)

func checkLetters(in string) bool {
	hasLetter := false
	for _, r := range in {
		if u.IsLetter(r) {
			hasLetter = true
			break
		}
	}

	if hasLetter {
		return true
	} else {
		return false
	}
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	fmt.Println("Input Remark:", remark)

    if s.HasSuffix(remark, "?") && remark == s.ToUpper(remark) && checkLetters(remark) {
        return "Calm down, I know what I'm doing!"
    }

    if remark == s.ToUpper(remark) && checkLetters(remark) {
        return "Whoa, chill out!"
    }

    if s.HasSuffix(s.TrimSpace(remark), "?") {
        return "Sure."
    }

    if s.TrimSpace(remark) == "" {
        return "Fine. Be that way!"
    }
    
	return "Whatever."
}
