package utils

import "time"

func TimeNow() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}

func ReadableTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
