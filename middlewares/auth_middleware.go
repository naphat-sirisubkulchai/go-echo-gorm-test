package middlewares

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
	"os"
)

func JWTMiddleware() echo.MiddlewareFunc {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
    return middleware.JWTWithConfig(middleware.JWTConfig{
        SigningKey: []byte(jwtSecret),
    })
}
