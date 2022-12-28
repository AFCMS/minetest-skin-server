package utils

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// TODO: use a regexp instead of mail.ParseAddress

func IsValidEmail(email string) bool {
	m, err := mail.ParseAddress(email)
	return err == nil && m.Name == ""
}
