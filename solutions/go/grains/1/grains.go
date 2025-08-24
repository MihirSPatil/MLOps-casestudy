package grains
import ("errors")

func Square(number int) (uint64, error) {
	if number < 1 || number > 64{
        return uint64(0), errors.New("Wrong input, should be between 0 and 64")
    }
    res := uint64(1)
    for i := 1; i < number; i++{
        res *= 2
    }
    return res, nil
}

func Total() uint64 {
    var sum uint64 = 0
    for i := 1; i <= 64; i++{
		val, _ := Square(i)
        sum += val
    }
    return sum
}
