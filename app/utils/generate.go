package utils

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func GenerateOtp() string {
	rand.Seed(uint64(time.Now().Unix()))
	otpInt := rand.Intn(999999-100000) + 100000
	return fmt.Sprintf("%06d", otpInt)
}
