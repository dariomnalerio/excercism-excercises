package expenses

import (
	"fmt"
)

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
	var filteredRecords []Record
	for _, v := range in {
		isTrue := predicate(v)
		if isTrue {
			filteredRecords = append(filteredRecords, v)
		}
	}
	return filteredRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	recordsByPeriod := Filter(in, ByDaysPeriod(p))
	var expenses float64
	for _, v := range recordsByPeriod {
		expenses += v.Amount
	}
	return expenses
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var categoryExists bool
	var expenses float64

	for _, v := range in {
		if v.Category == c {
			categoryExists = true
			break
		}
	}
	if !categoryExists {
		return expenses, fmt.Errorf("unkown category %s", c)

	}

	recordsByPeriod := Filter(in, ByDaysPeriod(p))
	recordsByPeriodAndCategory := Filter(recordsByPeriod, ByCategory(c))
	for _, v := range recordsByPeriodAndCategory {
		expenses += v.Amount
	}

	return expenses, nil
}
