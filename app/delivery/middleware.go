package delivery

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zakirkun/kas-ku/config"
)

func VerifyAccess() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.GetString("JWT_SECRET")),
	})
}
