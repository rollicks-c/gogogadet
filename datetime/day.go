package datetime

import "time"

func BeginOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func EndOfDay(date time.Time) time.Time {
	d := BeginOfDay(date)
	d = d.Add(time.Hour * 24)
	d = d.Add(time.Nanosecond * -1)
	return d
}
