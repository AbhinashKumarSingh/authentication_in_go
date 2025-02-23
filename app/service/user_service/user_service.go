package user_service

import (
	"errors"
	"time"

	refresh_token_repo "example.com/m/v2/db/refresh_token"
	"example.com/m/v2/db/user_repo"
	"example.com/m/v2/models"
	"example.com/m/v2/requests"
	"example.com/m/v2/response"
	"example.com/m/v2/service/token_service"
	"github.com/phuslu/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SignupUserService(input requests.UserReq) (*response.TokenResponse, error) {

	emailExist, err := user_repo.EmailExist(input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Msgf("CreateUserService-> error in checking email exixt for email: %s", input.Email)
		return nil, err
	}

	if emailExist.Email != "" {
		log.Error().Err(err).Msgf("CreateUserService-> email already exixt for email: %s", input.Email)
		return nil, errors.New("email already exist")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	userCreate := &models.Users{
		Email:        input.Email,
		PasswordHash: string(hash),
	}
	if err := user_repo.CreateUSer(userCreate); err != nil {
		log.Error().Err(err).Msgf("CreateUSer-> Error creating user for email: %s", input.Email)
		return nil, err
	}

	accessToken, _, err := token_service.GenerateJWT(userCreate.ID, 15*time.Minute, false)
	if err != nil {
		log.Error().Err(err).Msgf("generateJWT-> Error creating user token  for email: %s", input.Email)
		return nil, err
	}

	// Generate refresh token (expires in 7 days)
	refreshToken, refreshExp, err := token_service.GenerateJWT(userCreate.ID, 7*24*time.Hour, true)
	if err != nil {
		log.Error().Err(err).Msgf("generateJWT-> Error generating user refresh token  for email: %s", input.Email)
		return nil, err
	}

	refreshTokenRecord := &models.RefreshTokens{
		UserID:    userCreate.ID,
		Token:     refreshToken,
		ExpiresAt: refreshExp,
		Revoked:   false,
	}
	if err := refresh_token_repo.Create(refreshTokenRecord); err != nil {
		log.Error().Err(err).Msgf("generateJWT-> Error creating user refresh token  for email: %s", input.Email)
		return nil, err
	}
	res := &response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil

}

func SignInUserService(input requests.UserReq) (*response.TokenResponse, error) {

	emailExist, err := user_repo.EmailExist(input.Email)
	if err != nil {
		log.Error().Err(err).Msgf("CreateUserService-> error in checking email exixt for email: %s", input.Email)
		return nil, err
	}

	if emailExist.Email == "" {
		log.Error().Err(err).Msgf("CreateUserService-> account deos not exist for email: %s", input.Email)
		return nil, errors.New("email already exist")
	}

	if !comparePasswords(emailExist.PasswordHash, input.Password) {

		log.Error().Err(err).Msgf("CreateUserService-> password deos not match for email: %s", input.Email)
		return nil, errors.New("invalid credentials")
	}

	accessToken, _, err := token_service.GenerateJWT(emailExist.ID, 15*time.Minute, false)
	if err != nil {
		log.Error().Err(err).Msgf("generateJWT-> Error creating user token  for email: %s", input.Email)
		return nil, err
	}

	// Generate refresh token (expires in 7 days)
	refreshToken, refreshExp, err := token_service.GenerateJWT(emailExist.ID, 7*24*time.Hour, true)
	if err != nil {
		log.Error().Err(err).Msgf("generateJWT-> Error generating user refresh token  for email: %s", input.Email)
		return nil, err
	}

	refreshTokenRecord := &models.RefreshTokens{
		UserID:    emailExist.ID,
		Token:     refreshToken,
		ExpiresAt: refreshExp,
		Revoked:   false,
	}
	if err := refresh_token_repo.Create(refreshTokenRecord); err != nil {
		log.Error().Err(err).Msgf("generateJWT-> Error creating user refresh token  for email: %s", input.Email)
		return nil, err
	}
	res := &response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil
}

func comparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
