package armstrong
import ("fmt"
        "math"
       	"strconv")
func IsNumber(n int) bool {
    s := strconv.Itoa(n)
    digits := len(s)
    fmt.Println(digits)
    sum := 0
    for _, val := range s{
        num := int(val-'0')
        sum+= int(math.Pow(float64(num), float64(digits)))
    }
    if sum - n ==0{return true}else {return false}
}
