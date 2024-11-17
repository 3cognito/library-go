package utils

import (
	"fmt"
	"strconv"
	"time"
)

func TimeNow() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}

func ReadableTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ParseAccessTokenExpiryDuration(accessTokenExpiryDuration string) time.Duration {
	num, err := strconv.Atoi(accessTokenExpiryDuration)
	if err != nil {
		fmt.Println("Access token validity duration should be a valid number: ", err)
		panic(err)
	}

	return time.Hour * time.Duration(num)
}

func ParseStringTime(timeStamp string) (time.Time, error) {
	var err error
	var data time.Time

	formats := []string{"2006-01-02 15:04:05", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05+07:00"}

	for _, f := range formats {
		data, err = time.Parse(f, timeStamp)
		if err == nil {
			return data, nil
		}
	}

	return data, err
}

func ParseStringDate(timeStamp string) (time.Time, error) {
	return time.Parse("2006-01-02", timeStamp)
}
