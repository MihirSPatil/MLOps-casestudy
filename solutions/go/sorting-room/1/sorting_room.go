package sorting
import ("fmt"
       "strconv"
       )
// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumberBox interface {
	Value() string
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}


// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    switch typedef:= fnb.(type) {
        case FancyNumber:
        	val, _ := strconv.Atoi(typedef.Value())
        	return val
        default:
        	return 0
    }
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	int_val := ExtractFancyNumber(fnb)
    return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(int_val))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch typedef := i.(type){
        case int:
        	return DescribeNumber(float64(typedef))
        case float64:
        	return DescribeNumber(typedef)
        case NumberBox:
        	return DescribeNumberBox(typedef)
        case FancyNumberBox:
        	return DescribeFancyNumberBox(typedef)
        default :
        	return fmt.Sprint("Return to sender")
    }
}
