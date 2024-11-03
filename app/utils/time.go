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

func ParseAccessTokenExpiryTime(accessTokenExpiryDuration string) time.Time {
	num, err := strconv.Atoi(accessTokenExpiryDuration)
	if err != nil {
		fmt.Println("Access token validity duration should be a valid number: ", err)
		panic(err)
	}

	return time.Now().Add(time.Hour * time.Duration(num))
}
