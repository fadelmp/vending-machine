package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SwaggerRoute(routes *echo.Echo) {

	routes.GET("/swagger/*", echoSwagger.WrapHandler)
}
