package booking

import (
	"fmt"
	"time"
)

const (
	layoutNumeric = "1/2/2006 15:04:05"
	layoutLong    = "January 2, 2006 15:04:05"
	layoutFull    = "Monday, January 2, 2006 15:04:05"
)

func parseLayout(layout, date string) time.Time {
	t, _ := time.Parse(layout, date)
	return t
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return parseLayout(layoutNumeric, date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t := parseLayout(layoutLong, date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t := parseLayout(layoutFull, date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := parseLayout(layoutNumeric, date)
	formatted := t.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", formatted)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}