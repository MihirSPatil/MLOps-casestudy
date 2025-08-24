package pangram
import ("unicode")

func IsPangram(input string) bool {
    pangram := make(map[rune]bool)
	for _, char := range input{
        if unicode.IsLetter(char){
            pangram[unicode.ToLower(char)] = true
        }
    }
    return len(pangram) == 26
}
