package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	now := time.Now()
	layout := "January 2, 2006 15:04:05"
	appointment, _ := time.Parse(layout, date)
	return appointment.Before(now)

}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	hour := t.Hour()

	if hour >= 12 && hour <= 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	appointment, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %s", appointment.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	aniversary, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return aniversary
}
