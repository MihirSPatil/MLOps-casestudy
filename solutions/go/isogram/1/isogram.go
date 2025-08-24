package isogram
import "unicode"

func IsIsogram(word string) bool {
    seen := make(map[rune]bool)
    for _, v :=range(word){
        if !unicode.IsLetter(v){
            continue
        }
        v = unicode.ToLower(v)
        if seen[v] == true{
            return false
        }
        seen[v] = true
    }
    return true
}
