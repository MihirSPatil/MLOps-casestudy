package rotationalcipher
import ("strings"
        "unicode"
        )
func RotationalCipher(plain string, shiftKey int) string {
    plain = strings.TrimSpace(plain)
    var cipher []rune
    for _, val := range plain{
        if unicode.IsLower(val){
        shifted := ((val-'a')+rune(shiftKey))%26 + 'a'
        cipher = append(cipher, shifted)
            }else if unicode.IsUpper(val){
        shifted := ((val-'A')+rune(shiftKey))%26 + 'A'
        cipher = append(cipher, shifted)
            } else {cipher = append(cipher, val)}
        }
    return string(cipher)
}
