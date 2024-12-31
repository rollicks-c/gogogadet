package datetime

import "time"

func DateOnly(t time.Time) time.Time {
	y, m, d := t.Date()
	date := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	return date
}

func CreateTime(date time.Time, timeExp string) (time.Time, error) {

	// parse time
	timeOnly, err := time.Parse("15:04", timeExp)
	if err != nil {
		return time.Time{}, err
	}

	// create timestamp
	startTime := time.Date(date.Year(), date.Month(), date.Day(), timeOnly.Hour(), timeOnly.Minute(), 0, 0, date.Location())
	return startTime, nil
}
