package apis

import (
	"net/http"

	"example.com/m/v2/constants"
	"example.com/m/v2/requests"
	"example.com/m/v2/service/token_service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
)

func ValidateRefreshToken(c echo.Context) error {

	var input requests.RefreshTokenReq
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	tokenDetails, err := token_service.ValidateRefreshToken(input)
	if err != nil {
		log.Error().Err(err).Msg("ValidateRefreshToken: Failed to validate refresh token")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not validate refresh token"})
	}

	return c.JSON(http.StatusOK, tokenDetails)
}

func CheckToken(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, "Missing token")
	}

	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return constants.JwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, "Invalid or expired token")
	}

	return c.JSON(http.StatusOK, "Token is valid")
}

func RevokeToken(c echo.Context) error {
	var input requests.RefreshTokenReq
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	if err := token_service.RevokeToken(input); err != nil {
		log.Error().Err(err).Msg("RevokeToken: Failed to revoke token")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to revoke token"})
	}

	return c.JSON(http.StatusOK, "Token revoked successfully")
}
