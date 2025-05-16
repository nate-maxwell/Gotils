package time_utils

import (
	"time"
)

// Returns string: 'yyyymmdd'.
func GetDate() string {
	return time.Now().Format("20060102")
}

// Returns string: 'HH:MM:SS:XX', X is microsecond.
func GetTime() string {
	return time.Now().Format("15:04:05:00")
}

// Whether the given timestamp is between times a and b.
func IsTimeBetweenAAndB(curTime, a, b time.Time) bool {
	if curTime.After(a) && curTime.Before(b) {
		return true
	}
	return false
}

func GetTimestampeOfTimezone(timeZone string) (time.Time, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return time.Now(), err
	}
	curTime := time.Now().In(loc)
	return curTime, nil
}

// Returns the number of days between two given timestamps.
func DiffInDays(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

// Returns the number of specific weekdays between two timestamps.
// 1 = Monday, 2 = Tuesday ... 7 = Sunday.
func NumWeekdaysBetween(day int, start, end time.Time) int {
	totalDays := 0
	for start.Before(end) {
		if int(start.Weekday()) == day {
			totalDays += 1
		}
		start = start.AddDate(0, 0, 1)
	}
	return totalDays
}

// Is the given year a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
