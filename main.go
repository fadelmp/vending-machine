// You can edit this code!
// Click here and start typing.
package main

import (
	"os"
	config "vending-machine/config"
	routes2 "vending-machine/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.ReadEnv()
	route := echo.New()

	dbConfig := config.InitDB()
	redisConfig := config.InitRedis()

	routes := routes2.Init(route, dbConfig, redisConfig)

	// set logger
	routes.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, time:${time_unix}, uri=${uri}, status=${status}, error=${error}, latency_human=${latency}, bytes_in=${bytes_in}, bytes_out=${bytes_out} \n",
	}))

	// Gzip Compression
	routes.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	service_port := os.Getenv("SERVICE_PORT")
	routes.Logger.Fatal(routes.Start(service_port))

}
