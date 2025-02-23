package apis

import (
	"net/http"

	"example.com/m/v2/requests"
	"example.com/m/v2/service/user_service"
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
)

func UserSignup(c echo.Context) error {

	var userInput requests.UserReq
	if err := c.Bind(&userInput); err != nil {
		log.Error().Err(err).Msg("UserSignup: Failed to bind request")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	tokenDetails, err := user_service.SignupUserService(userInput)
	if err != nil {
		log.Error().Err(err).Msg("UserSignup: Failed to signup user")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
	}

	return c.JSON(http.StatusOK, tokenDetails)
}

func UserSignin(c echo.Context) error {

	var userInput requests.UserReq
	if err := c.Bind(&userInput); err != nil {
		log.Error().Err(err).Msg("UserSignin: Failed to bind request")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	tokenDetails, err := user_service.SignInUserService(userInput)
	if err != nil {
		log.Error().Err(err).Msg("UserSignin: Failed to signin user")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not signin user"})
	}

	return c.JSON(http.StatusOK, tokenDetails)
}
