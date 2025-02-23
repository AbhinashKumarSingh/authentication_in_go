package main

import (
	"example.com/m/v2/apis"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo) {
	v2 := e.Group("/user-service")
	userRoutes(v2)
}

func userRoutes(router *echo.Group) {
	routes := router.Group("/user")
	routes.POST("/signup", apis.UserSignup)
	routes.POST("/signin", apis.UserSignin)
	routes.POST("/refresh-token", apis.ValidateRefreshToken)
	routes.POST("/revoke-token", apis.RevokeToken)
}
