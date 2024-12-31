package datetime

import "time"

func BeginOfWeek(date time.Time) time.Time {
	weekDay := date.Weekday()
	offset := time.Duration(0)
	if weekDay == time.Sunday {
		offset = time.Duration(time.Saturday)
	} else {
		offset = time.Duration(weekDay - 1)
	}
	return date.Add(time.Hour * time.Duration(24) * -offset)

}

func EndOfWeek(date time.Time) time.Time {
	weekDay := date.Weekday()
	offset := time.Duration(0)
	if weekDay == time.Sunday {
		offset = time.Duration(0)
	} else {
		offset = time.Duration(7 - weekDay)
	}
	return date.Add(time.Hour * time.Duration(24) * offset)

}
