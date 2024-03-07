package database

import "minetest-skin-server/models"

// AccountList Return users that are not banned
func AccountList() ([]models.Account, error) {
	var result []models.Account

	if err := DB.Find(&result).Where("ban = false").Error; err != nil {
		return nil, err
	}

	return result, nil
}
