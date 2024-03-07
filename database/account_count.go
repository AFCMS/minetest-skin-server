package database

import "minetest-skin-server/models"

// AccountCount Count accounts registered
func AccountCount() (int64, error) {
	var count int64 = 0
	if err := DB.Model(&models.Account{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}
