package wlog

import (
	"time"
)

func Week(t time.Time) int {
	switch t.Weekday() {
	case time.Monday:
		return 1
	case time.Tuesday:
		return 2
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 5
	case time.Saturday:
		return 6
	case time.Sunday:
		return 7
	}
	return 0
}

func WeekString(t time.Time) string {
	return t.Weekday().String()
}


func WeekStringShort(t time.Time) string {
	s := t.Weekday().String()
	return string([]byte(s)[:3])
}

func Month(t time.Time) int {
	switch t.Month() {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	case time.December:
		return 12
	}
	return 0
}

func MonthString(t time.Time) string {
	return t.Month().String()
}

func MonthStringShort(t time.Time) string {
	s := t.Month().String()
	return string([]byte(s)[:3]) 
}

func Year(t time.Time) int {
	return t.Year()
}

func YearShort(t time.Time) int {
	return t.Year()%100
}

func IsAm(t time.Time) bool {
	return (t.Hour() < 12) || (t.Hour() == 24) 
}

func AmPmhour(t time.Time) int {
	return t.Hour() % 12
}


func formatTime(t time.Time) string {
    return t.Format("2006-01-02 15:04:05")
}