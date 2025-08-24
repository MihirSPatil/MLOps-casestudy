package reverse

func Reverse(input string) string {
	reverse := ""
    runes := []rune(input)
    for i := len(runes)-1; i >=0; i--{
        reverse += string(runes[i])
    }
    return reverse
}
