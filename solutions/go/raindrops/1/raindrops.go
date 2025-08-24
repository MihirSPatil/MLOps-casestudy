/*
Raindrops is a slightly more complex version of the FizzBuzz challenge, 
given a number convert it to its corresponding raindrop sounds based on 
whether it is divisible by 3,5 or 7
*/
package raindrops
import "strconv"

func Convert(number int) string {
	var res_str string
    var divisors = []int{3,5,7}
    sounds := []string{"Pling", "Plang", "Plong"}
    for idx, div := range divisors{
        if number % div == 0{
            res_str = res_str + sounds[idx]
        }
    }
    if res_str != ""{
        return res_str
    } else {
        return strconv.Itoa(number)
    }
}
