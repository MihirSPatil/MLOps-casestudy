package prime
import ("errors"
		"math"
		)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
    if n <= 0{
        return 0, errors.New("input 'n' has to be greater than zero")
    }
    
    var primeNo []int
    i:=2
    for len(primeNo) < n{
        factors:= int(math.Sqrt(float64(i)))
        isPrime := true
        for j:=2; j <= factors; j++{
            if i%j == 0{
                isPrime = false
                break
            }
        }
        if isPrime{
            primeNo = append(primeNo, i)
        }
        i++
    }
    return primeNo[n-1], nil
}
