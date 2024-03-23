package routes

import (
	"vending-machine/injection"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Init(routes *echo.Echo, db *gorm.DB, redis *redis.Client) *echo.Echo {

	// Vending Route & Injection
	vending := injection.VendingInjection(db, redis)
	VendingRoute(routes, vending)

	return routes
}
