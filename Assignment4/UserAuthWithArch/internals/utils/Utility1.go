package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"crypto/sha256"
)


func GenerateOtp() string {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		panic("Issue with Otp generation")
	}
	return fmt.Sprintf("%06d", (int(b[0])<<16|int(b[1])<<8|int(b[2]))%1000000)

}
func PasswordEncrpytion(pwd string) string{
	hash := sha256.New()
	hash.Write([]byte(pwd))
	return hex.EncodeToString(hash.Sum(nil))

}