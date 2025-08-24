package resistorcolorduo
import "strings"
import "strconv"
// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
    col_hash := map[string]string{
        "black":"0",
        "brown":"1",
        "red": "2",
        "orange": "3",
        "yellow": "4",
        "green": "5",
        "blue": "6",
        "violet": "7",
        "grey": "8",
        "white": "9",
    }
    sum := ""
    for i:=0; i < 2; i++{
        sum += col_hash[strings.ToLower(colors[i])]
    }
    num, _ := strconv.Atoi(sum)
    return num
}
