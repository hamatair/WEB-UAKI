package middlewares

import (
	"backend_server/jwt"
	"errors"
	"strings"

	"github.com/gobuffalo/buffalo"
)

func AuthMiddleware(jwtService jwt.Interface) buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.Error(401, errors.New("missing Authorization header"))
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.Error(400, errors.New("invalid Authorization format, use 'Bearer <token>'"))
			}

			adminId, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				return c.Error(401, errors.New("invalid or expired token"))
			}

			c.Set("admin_id", adminId.String())
			return next(c)
		}
	}
}
