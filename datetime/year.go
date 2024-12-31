package datetime

import "time"

func BeginOfYear(date time.Time) time.Time {
	y, _, _ := date.Date()
	return time.Date(y, 1, 1, 0, 0, 0, 0, date.Location())
}

func EndOfYear(date time.Time) time.Time {
	y, _, _ := date.Date()
	nextYear := time.Date(y+1, 1, 1, 0, 0, 0, 0, date.Location())
	return nextYear.Add(time.Second * -1)
}
