package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

func GenerateOtp() int {
	rand.Seed(uint64(time.Now().Unix()))
	return rand.Intn(999999)
}
