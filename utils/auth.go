package utils

import (
	"errors"
	"minetest-skin-server/database"
	"minetest-skin-server/models"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*models.Account, error) {
	db := database.DB
	var account models.Account
	if err := db.Where(&models.Account{Email: e}).Find(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func getUserByUsername(u string) (*models.Account, error) {
	db := database.DB
	var user models.Account
	if err := db.Where(&models.Account{Name: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
