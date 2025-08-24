// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// https://golang.org/doc/effective_go.html#commentary
// This package takes in an input of type time adds a gigasecond 
// and returns the resulting value in a human readable format
package gigasecond

// import path for the time package from the standard library
import ("time")

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
    return t.Add(time.Second * 1e9)
}
