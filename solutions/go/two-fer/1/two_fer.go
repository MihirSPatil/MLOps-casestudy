// Package twofer provides a function to generate the "twofer" phrase.
// The twofer phrase is "One for you, one for me."

package twofer
import "fmt"
// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
    var name_str string
    switch name{
        case "Alice":
            name_str = "Alice"
        case "Bob":
        	name_str = "Bob"
        default :
        	name_str = "you"
    }
	return fmt.Sprintf("One for %s, one for me.", name_str)
}
