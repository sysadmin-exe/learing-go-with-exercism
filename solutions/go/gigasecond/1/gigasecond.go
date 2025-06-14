// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"
import "fmt"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	f := t.Format("2006-01-02T15:04:05")
	fmt.Println(f)
	
	p, err := time.Parse("2006-01-02T15:04:05", f)
	if err != nil {
		panic(err)
	}

	ags := p.Add(1000000000 * time.Second)
	fmt.Println(ags)
	return ags
}