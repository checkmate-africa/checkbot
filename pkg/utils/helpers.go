package utils

import (
	"strconv"
	"time"
)

type DmyDate struct {
	Day      string
	Month    string
	Year     string
	DaysTime string
}

func GetNextDayOfWeek(dayOfWeek int) (day string, month string, year string, daysUntil string) {
	var (
		CurrentDate   = time.Now()
		DaysUntilNext = (dayOfWeek + 7 - int(CurrentDate.Weekday())) % 7
		NextDate      = CurrentDate.AddDate(0, 0, DaysUntilNext)
	)

	y, m, d := NextDate.Date()

	return strconv.Itoa(d), m.String(), strconv.Itoa(y), strconv.Itoa(DaysUntilNext)
}
