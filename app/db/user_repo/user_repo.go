package user_repo

import (
	"example.com/m/v2/config"
	"example.com/m/v2/models"
)

func CreateUSer(user *models.Users) error {

	return config.DB.
		Model(&models.Users{}).
		Create(user).Error

}

func EmailExist(email string) (*models.Users, error) {

	user := new(models.Users)
	err := config.DB.
		Model(&models.Users{}).
		Where("email = ?", email).
		First(&user).Error

	return user, err
}
