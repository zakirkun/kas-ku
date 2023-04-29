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

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
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

	app := v1.Group("/app")
	{
		app.Use(delivery.VerifyAccess())
		app.POST("/set-pin", userDelivery.SetPIN)
	}
}

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"message": "Services up"})
}
