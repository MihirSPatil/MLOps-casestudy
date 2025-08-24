// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

//Given an integer representation of a year this function determines if it is a leap year or not
package leap

// Accepts and integer and returns a bool value indicating if the year is a leap year or not
func IsLeapYear(year int) bool {
	if year %4 == 0{
        if year % 100 == 0{
        	if year % 400 == 0{
            	return true
        	}else {return false}
    	} else {return true}
    }else{return false}
}
