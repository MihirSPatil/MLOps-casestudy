package isbn
import ("fmt"; "strings"; "unicode"; "unicode/utf8")

func IsValidISBN(isbn string) bool {
    sum := 0
    var digit int
    cleanIsbn:= strings.Replace(isbn, "-", "", -1)
    if utf8.RuneCountInString(cleanIsbn) == 10{
    for idx, val := range cleanIsbn{
        if unicode.IsDigit(val){
            digit = int(val-'0')
        }else if unicode.IsLetter(val) && val == 'X' && idx == 9{
            digit = 10
        }else {return false}
        sum += (10-idx) * digit
        fmt.Printf("%d: %d, %d\n", (10-idx), digit, sum)
    }
    return sum % 11 == 0
    }else {return false}
}
