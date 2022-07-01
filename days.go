package utils

import "time"

func DaySlice(start, end time.Time) (DaySlice []time.Time) {
	currDate := start

	for end.After(currDate) == true {
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

func PrevMonthFirstLast() (time.Time, time.Time) {
	m := PreviousMonth()
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
