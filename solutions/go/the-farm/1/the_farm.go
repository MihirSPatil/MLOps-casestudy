package thefarm
import ("fmt"; "errors")
// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, num_cows int) (float64, error){
    total_fodder_available, fodder_err := fc.FodderAmount(num_cows)
    fattening_factor, fat_err := fc.FatteningFactor()
    if fodder_err != nil {
        return 0, fodder_err
    } else if fat_err != nil{
        return 0, fat_err
    } else {
        food_per_cow := (total_fodder_available * fattening_factor) / float64(num_cows)
        return food_per_cow, nil
    }
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, num_cows int)(float64, error){
    var res float64
    var err error
    if num_cows >0{
        res, err = DivideFood(fc, num_cows)
    } else {
        err = errors.New("invalid number of cows")
    }
    return res, err
}
type InvalidCowsError struct{
    num_cows int
    err_msg string
}

func (ice *InvalidCowsError) Error() string{
    return fmt.Sprintf("%d cows are invalid: %s",ice.num_cows, ice.err_msg)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(num_cows int)(error){
    if num_cows < 0{
        return &InvalidCowsError{
            num_cows: num_cows,
            err_msg: "there are no negative cows",
        }
    }else if num_cows == 0{
        return &InvalidCowsError{
            num_cows: num_cows,
            err_msg: "no cows don't need food",
        }
    }else {
        return nil
    }
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
