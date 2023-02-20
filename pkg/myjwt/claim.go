package myjwt

import (
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	jwt.StandardClaims
	UserID  uint `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}

var (
	ContextKey = "user"
)

func Middleware(key string) echo.MiddlewareFunc {
	conf := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(Claims)
		},
		SigningKey:  []byte(key),
		ContextKey:  ContextKey,
		TokenLookup: "header:Authorization:jwt",
	}

	return echojwt.WithConfig(conf)
}

func GetClaims(c echo.Context) *Claims {
	user := c.Get(ContextKey).(*jwt.Token)

	if !user.Valid {
		return nil
	}

	claims := user.Claims.(*Claims)

	return claims
}

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := GetClaims(c)
			if claims == nil || claims.IsAdmin == false {
				return echo.ErrForbidden
			}

			return next(c)
		}
	}
}
