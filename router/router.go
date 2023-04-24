package router

import (
	"net/http"

	"github.com/zakirkun/kas-ku/app/delivery"
	"github.com/zakirkun/kas-ku/app/repository"
	"github.com/zakirkun/kas-ku/app/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func RegisterRouter(e *echo.Echo, db *gorm.DB) {

	// middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// repository
	userRepository := repository.NewUsersRepository(db)

	// services
	userServices := services.NewUsersServices(userRepository)

	// delivery
	userDelivery := delivery.NewUsersDelivery(userServices)

	// routing
	e.GET("/", Ping)

	// grouping
	v1 := e.Group("/v1")

	v1.POST("/register", userDelivery.Register)
	v1.POST("/activation", userDelivery.Activation)
}

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"message": "Services up"})
}
