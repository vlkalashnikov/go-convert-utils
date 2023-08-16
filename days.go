package utils

import (
	"strconv"
	"time"
)

func DaySlice(start, end time.Time) (DaySlice []time.Time) {
	currDate := start

	for end.After(currDate) {
		DaySlice = append(DaySlice, currDate)
		currDate = currDate.AddDate(0, 0, int(1))
	}
	DaySlice = append(DaySlice, end)

	return
}

func FilterDaySlice(ss []time.Time, filter func(time.Time) bool) (res []time.Time) {
	for _, s := range ss {
		if filter(s) {
			res = append(res, s)
		}
	}
	return
}

func NotWeekend(day time.Time) bool {
	weekday := day.Weekday().String()
	return weekday != "Sunday" && weekday != "Saturday"
}

func IsWeekend(day time.Time) bool {
	weekday := day.Weekday().String()
	return weekday == "Sunday" || weekday == "Saturday"
}

func PreviousMonth() time.Time {
	currentTime := time.Now()
	return currentTime.AddDate(0, -1, 0)
}

func NextMonth() time.Time {
	currentTime := time.Now()
	return currentTime.AddDate(0, 1, 0)
}

func PrevMonthFirstLast() (time.Time, time.Time) {
	m := PreviousMonth()
	return BeginningOfMonth(m), EndOfMonth(m)
}

func NextMonthFirstLast() (time.Time, time.Time) {
	m := NextMonth()
	return BeginningOfMonth(m), EndOfMonth(m)
}

func BeginningOfMonth(d time.Time) time.Time {
	return d.AddDate(0, 0, -d.Day()+1)
}

func EndOfMonth(d time.Time) time.Time {
	return d.AddDate(0, 1, -d.Day())
}

func GetYearMonthDay(d time.Time) (year, month, day int) {
	year, mnth, day := d.Date()
	month = int(mnth)
	return
}

func GetYear(d time.Time) (year int) {
	year, _, _ = d.Date()
	return
}

func GetMonth(d time.Time) int {
	_, mnth, _ := d.Date()
	return int(mnth)
}

func GetDay(d time.Time) (day int) {
	_, _, day = d.Date()
	return
}

func GetCurrentYearMonthDay() (year, month, day int) {
	return GetYearMonthDay(time.Now())
}

func GetCurrentYear() int {
	return GetYear(time.Now())
}

func GetCurrentDay() int {
	return GetDay(time.Now())
}

func GetCurrentMonth() int {
	return GetMonth(time.Now())
}

func Month2Str(mnth int) string {
	if mnth < 10 {
		return "0" + strconv.Itoa(mnth)
	}
	return strconv.Itoa(mnth)
}

func Diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func DateBetween2Dates(date, start, end time.Time) bool {
	return (start.Before(date) || start.Equal(date)) && (end.After(date) || end.Equal(date))
}

func IsLastDayOfMonth(date time.Time) bool {
	return date.Day() == EndOfMonth(date).Day()
}

func GetNearestLeftDate(dates []time.Time, currentDate time.Time) time.Time {
	var minDiff int64 = -1
	var minDate time.Time

	for _, date := range dates {
		if date.Before(currentDate) {
			diff := currentDate.Unix() - date.Unix()
			if (minDiff == -1) || (diff < minDiff) {
				minDiff = diff
				minDate = date
			}
		}
	}
	return minDate
}

func ValidDate(dateStr time.Time) bool {
	if dateStr.IsZero() {
		return false
	}
	if GetYear(dateStr) == 1970 {
		return false
	}
	return true
}
