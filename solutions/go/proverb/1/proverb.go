//Package creates a rhyme based on the input words given
package proverb
import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
    var song []string
    rhymeLen := len(rhyme)
    if rhymeLen == 0 {
        return song
    }
    for i:=0; i < rhymeLen-1; i++ {
        song = append(song, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
    }
    song = append(song, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
    return song
}
