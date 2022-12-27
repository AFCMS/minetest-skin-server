package database

import "minetest-skin-server/models"

// Get account from ID
func AccountFromID(id int) (models.Account, error) {
	var a = models.Account{}

	if err := DB.Where("id = ?", id).First(&a).Error; err != nil {
		return a, err
	}
	return a, nil
}
