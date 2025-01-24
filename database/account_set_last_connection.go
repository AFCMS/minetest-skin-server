package database

import (
	"luanti-skin-server/models"
	"time"
)

func AccountSetLastConnection(a *models.Account) error {
	//a.LastConnection = time.Now()

	if err := DB.Model(&a).Update("last_connection", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
