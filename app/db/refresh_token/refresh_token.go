package refresh_token_repo

import (
	"example.com/m/v2/config"
	"example.com/m/v2/models"
)

func Create(create *models.RefreshTokens) error {

	return config.DB.
		Model(&models.RefreshTokens{}).
		Create(create).Error

}

func Updates(userID int64, updates map[string]interface{}) error {
	return config.DB.
		Model(&models.RefreshTokens{}).
		Where("user_id = ?", userID).
		Updates(updates).Error

}
func UpdateByRefreshToken(refreshToken string, updates map[string]interface{}) error {
	return config.DB.
		Model(&models.RefreshTokens{}).
		Where("token = ?", refreshToken).
		Updates(updates).Error

}
