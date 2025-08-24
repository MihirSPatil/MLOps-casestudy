package romannumerals
import ("strings"
        "strconv"
	   "errors")

var converter = []struct{
    digit int
    sym string
}{
    {1000,"M"},
    {900,"CM"},
    {500,"D"},
    {400,"CD"},
    {100,"C"},
    {90,"XC"},
    {50,"L"},
    {40,"XL"},
    {10,"X"},
    {9,"IX"},
    {5,"V"},
    {4,"IV"},
    {1,"I"},
}
func ToRomanNumeral(input int) (string, error) {
    if input <=0 || input >3999{
        return strconv.Itoa(input), errors.New("Input is not in the valid range")
    }

    var roman_num strings.Builder
    for _, val := range converter{
        for input >= val.digit{
            roman_num.WriteString(val.sym)
            input -= val.digit
        }
    }
    
    return roman_num.String(), nil
}
