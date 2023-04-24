package main

import (
	"flag"
	"fmt"

	"github.com/zakirkun/kas-ku/migration"

	"github.com/labstack/echo"
	"github.com/zakirkun/kas-ku/database"
	"github.com/zakirkun/kas-ku/router"

	"github.com/zakirkun/kas-ku/config"
	"github.com/zakirkun/kas-ku/logger"
)

var cfgFile *string

func init() {
	cfgFile = flag.String("c", "config.toml", "Configuration")
}

func main() {

	// init config
	if err := config.Initialize(*cfgFile); err != nil {
		// log.Fatalf("Error reading configuration: %s\n", err.Error())
		logger.Logger.Error().Str("ERROR", fmt.Sprintf("Error reading configuration: %s\n", err.Error()))
	}

	// init echo
	server := echo.New()

	// init database
	db := database.OpenDb()

	// run migrations
	migration.Migrate(db)

	// init router
	router.RegisterRouter(server, db)

	// init web server
	server.Logger.Fatal(server.Start(":" + config.GetString("PORT")))
}
