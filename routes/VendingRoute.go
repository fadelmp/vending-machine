package routes

import (
	"vending-machine/handler"

	"github.com/labstack/echo/v4"
)

func VendingRoute(routes *echo.Echo, handler *handler.VendingHandler) {

	vending := routes.Group("/vending")
	{
		vending.GET("", handler.GetAll)
		vending.GET("/:id", handler.GetById)

		vending.POST("", handler.Create)
		vending.PUT("", handler.Update)
		vending.DELETE("/:id", handler.Delete)
	}
}
