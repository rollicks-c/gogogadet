package datetime

import "time"

func BeginOfMonth(date time.Time) time.Time {
	y, m, _ := date.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, date.Location())
}

func EndOfMonth(date time.Time) time.Time {
	y, m, _ := date.Date()
	nextMonth := time.Date(y, m+1, 1, 0, 0, 0, 0, date.Location())
	return nextMonth.Add(time.Second * -1)
}
