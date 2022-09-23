package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		t, err = time.Parse("January 2, 2006 15:04:05", date)
	}
	if err != nil {
		t, err = time.Parse("Monday, January 2, 2006 15:04:05", date)
	}
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	hour, _, _ := Schedule(date).Clock()
	return hour > 11 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	d := Schedule(date)
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", 
		d.Weekday(), d.Month(), d.Day(), d.Year(), d.Hour(), d.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return Schedule(fmt.Sprintf("9/15/%v 00:00:00", time.Now().Year()))
}
