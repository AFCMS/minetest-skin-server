package database

import "luanti-skin-server/models"

func SkinApproval(s *models.Skin, state bool) error {
	if err := DB.Model(&s).Update("approved", state).Error; err != nil {
		return err
	}
	return nil
}
