package util

import (
	"github.com/raja/argon2pw"
)

func Encrypt(password string) (string, error) {
	return argon2pw.GenerateSaltedHash(password)
}

func Validate(encrypted, raw string) (bool, error) {
	return argon2pw.CompareHashWithPassword(encrypted, raw)
}
