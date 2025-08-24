package booking

import ("time"
        "fmt"
       )

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	formattedTime, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{} // Return zero-value time
	}
	return formattedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    appointment, err := time.Parse(layout, date)
    if err != nil {
        fmt.Println("Could not parse date", err)
        return false
    }
    return appointment.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    appointment, err := time.Parse(layout, date)
    if err!= nil{
        fmt.Println("Could not parse date", err)
        return false
    }
    return appointment.Hour() >= 12 && appointment.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    description, _ := time.Parse("1/2/2006 15:04:05", date)
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%02d.", description.Weekday(), description.Month(), description.Day(), description.Year(), description.Hour(), description.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	date :=time.Date(time.Now().Year(), time.September, 15, 0,0,0,0, time.UTC)
    return date
}
