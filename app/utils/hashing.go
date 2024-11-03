package utils

import "golang.org/x/crypto/bcrypt"

func HashData(data string) string {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(data), 14)
	if err != nil {
		panic(err)
	}
	return string(encrypted)
}

func VerifyDataHash(data, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err == nil
}
