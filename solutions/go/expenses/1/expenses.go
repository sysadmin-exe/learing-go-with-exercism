package expenses

import "errors"
import "fmt"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
    r := []Record{}
    for _, i := range in {
        if predicate(i) {
            r = append(r, i)
        }
    }
    return r
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
    return func(r Record) bool {
        if r.Day >= p.From && r.Day <= p.To {
            return true
        } else {
            return false
        }
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
    return func(r Record) bool {
        if r.Category == c {
            return true
        } else {
            return false
        }
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	amountSpent := 0.00
    records := Filter(in, ByDaysPeriod(p)) 
    for _, i := range records {
        amountSpent += i.Amount
    }
    return amountSpent
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.


func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	amountSpent := 0.00
    dayPeriodRecords := Filter(in, ByDaysPeriod(p)) 
    categoryRecords := Filter(dayPeriodRecords, ByCategory(c)) 
    if len(dayPeriodRecords) == 0  {
        return amountSpent, nil
    }
    if len(categoryRecords) == 0  {
        return amountSpent, errors.New(fmt.Sprintf("unknown category %s", c))
    }
    for _, i := range categoryRecords {
        amountSpent += i.Amount
    }
    return amountSpent, nil
}
