// Package bob, returns the standard responses of the most boring teenager named Bob
package bob
import (
    "strings"
    "unicode"
)

// Hey takes in a standard string as input for an interaction with Bob and then returns the response as a string
func Hey(remark string) string {
    remark = strings.TrimSpace(remark)
    var yells bool = false
    var hasLetter bool = false

    for _, letter := range remark {
        if unicode.IsLetter(letter) {
            hasLetter = true
            if !unicode.IsUpper(letter) {
                yells = false
                break
            }
            yells = true
        }
    }
    yells = yells && hasLetter

    switch {
        case remark == "":
            return "Fine. Be that way!"
        case yells && strings.HasSuffix(remark, "?"):
            return "Calm down, I know what I'm doing!"
        case yells:
            return "Whoa, chill out!"
        case strings.HasSuffix(remark, "?"):
            return "Sure."
        default:
            return "Whatever."
    }
}
