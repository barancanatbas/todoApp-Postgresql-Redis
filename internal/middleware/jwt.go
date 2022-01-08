package middleware

import (
	"todo/repository"

	config "todo/internal/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*config.JwtCustom)

		userlogin := claims.User
		// vt kontrolü
		_, err := repository.Get().User().Login(userlogin.UserName)
		if err != nil {
			return c.JSON(200, "Bilinmeyen bir hata oluştu")
		}

		return next(c)
	}
}
