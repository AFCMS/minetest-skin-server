package database

import "minetest-skin-server/models"

// AccountFromEmail Get account from ID
func AccountFromEmail(email string) (models.Account, error) {
	var a = models.Account{}

	if err := DB.Where("email = ?", email).First(&a).Error; err != nil {
		return a, err
	}
	return a, nil
}
