package database

import "luanti-skin-server/models"

// AccountDelete Delete account
func AccountDelete(name string) error {
	return DB.Where("name = ?", name).Delete(&models.Account{}).Error
}
