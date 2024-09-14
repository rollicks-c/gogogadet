package datetime

import "time"

func BeginOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func EndOfDay(date time.Time) time.Time {
	d := BeginOfDay(date)
	d = d.Add(time.Hour * 1)
	d = d.Add(time.Nanosecond - 1)
	return d
}

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

func BeginOfMonth(date time.Time) time.Time {
	y, m, _ := date.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, date.Location())
}

func EndOfMonth(date time.Time) time.Time {
	y, m, _ := date.Date()
	nextMonth := time.Date(y, m+1, 1, 0, 0, 0, 0, date.Location())
	return nextMonth.Add(time.Second * -1)
}

func BeginOfYear(date time.Time) time.Time {
	y, _, _ := date.Date()
	return time.Date(y, 1, 1, 0, 0, 0, 0, date.Location())
}

func EndOfYear(date time.Time) time.Time {
	y, _, _ := date.Date()
	nextYear := time.Date(y+1, 1, 1, 0, 0, 0, 0, date.Location())
	return nextYear.Add(time.Second * -1)
}

func DateOnly(t time.Time) time.Time {
	y, m, d := t.Date()
	date := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	return date
}
