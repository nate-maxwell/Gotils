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
