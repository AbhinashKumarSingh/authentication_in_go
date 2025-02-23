package token_service

import (
	"time"

	"example.com/m/v2/constants"
	refresh_token_repo "example.com/m/v2/db/refresh_token"
	"example.com/m/v2/requests"
	"example.com/m/v2/response"
	"github.com/golang-jwt/jwt"
	"github.com/phuslu/log"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID int64, duration time.Duration, isRefresh bool) (string, time.Time, error) {
	expirationTime := time.Now().Add(duration)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := constants.JwtSecretKey
	if isRefresh {
		secretKey = constants.JwtRefreshSecretKey
	}
	tokenString, err := token.SignedString(secretKey)
	return tokenString, expirationTime, err
}

func ValidateRefreshToken(refreshToken requests.RefreshTokenReq) (*response.TokenResponse, error) {

	// Validate refresh token
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(refreshToken.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JwtRefreshSecretKey), nil
	})
	// if err != nil || !token.Valid {
	// 	log.Error().Err(err).Msg("Invalid refresh token")
	// }
	// claims, ok := token.Claims.(*Claims)
	// if !ok || !token.Valid {
	// 	// return 0, errors.New("invalid token")
	// }

	// return 	userID, nil

	// Generate new access token

	var userID int64
	if id, ok := claims["user_id"]; ok {
		userID = id.(int64)
	}
	accessToken, _, err := GenerateJWT(userID, 15*time.Minute, false)
	if err != nil {
		log.Error().Err(err).Msgf("generateJWT-> Error creating user token  for userID: %d", userID)
		return nil, err
	}

	// Generate refresh token (expires in 7 days)
	newRefreshToken, refreshExp, err := GenerateJWT(userID, 7*24*time.Hour, true)
	if err != nil {
		log.Error().Err(err).Msgf("generateJWT-> Error generating user refresh token  for userID: %d", userID)
		return nil, err
	}
	updates := map[string]interface{}{
		"token":      newRefreshToken,
		"expires_at": refreshExp,
	}
	if err := refresh_token_repo.Updates(userID, updates); err != nil {
		log.Error().Err(err).Msgf("refresh_token_repo.Updates-> Error updating  refresh token  for userID: %d", userID)
		return nil, err
	}

	return &response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, err
}

func RevokeToken(refreshToken requests.RefreshTokenReq) error {
	updates := map[string]interface{}{
		"revoked": true,
	}

	if err := refresh_token_repo.UpdateByRefreshToken(refreshToken.RefreshToken, updates); err != nil {
		return err
	}
	return nil
}
