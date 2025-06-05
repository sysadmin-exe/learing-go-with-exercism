package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout, date)
    fmt.Println(t)
 	return t 
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
	schedule, _ := time.Parse(layout, date)
    now := time.Now()
    if now.After(schedule) {
        return true
    } else {
        return false
    }
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    fmt.Println(date)
    layout := "Monday, January 2, 2006 15:04:05"
	schedule, _ := time.Parse(layout, date)
	compareTime := time.Date(0, 1, 1, schedule.Hour(), schedule.Minute(), schedule.Second(), 0, time.UTC)
    noon, _ := time.Parse(time.TimeOnly, "12:00:00")
    evening,_ := time.Parse(time.TimeOnly, "18:00:00")
    fmt.Println(compareTime)
    if noon.Before(compareTime) && evening.After(compareTime) {
        fmt.Println(true)
        return true
    } else {
        fmt.Println(false)
        return false
    }
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    fmt.Println(date)
    layout := "1/2/2006 15:04:05"
    d, _ := time.Parse(layout, date)
    layoutFormated := "You have an appointment on Monday, January 2, 2006, at 15:04."
    formated := d.Format(layoutFormated)
    fmt.Println(formated)
    return formated
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    anniversary, _ := time.Parse("2006-02-01", "2025-15-09")
	return anniversary
}
