package main

import (
	"alta/immersive-dashboard-api/app/config"
	"alta/immersive-dashboard-api/app/database"
	"alta/immersive-dashboard-api/app/migration"
	"alta/immersive-dashboard-api/app/routers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.ReadEnv()
	database := database.InitDB(config)
	errMigrate := migration.InitMigrate(database)
	if errMigrate != nil {
		log.Fatal(errMigrate)
	}

	echo := echo.New()
	echo.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	echo.Pre(middleware.RemoveTrailingSlash())
	echo.Use(middleware.CORS())


	routers.InitRouters(database, echo)

	echo.Logger.Fatal(echo.Start(":8080"))
}
