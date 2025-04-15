package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yuhari7/backend_supervision/pkg/jwt"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing authorization header"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token format"})
		}

		tokenStr := parts[1]
		fmt.Println("Received Token:", tokenStr)

		claims, err := jwt.ParseAccessToken(tokenStr)
		fmt.Println("Token Claims:", claims)

		if err != nil {
			log.Println("Error Parsing Token:", err) // Debug log
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid or expired token"})
		}

		// simpan claims ke context
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role_id", claims.RoleID)

		return next(c)
	}
}

// Admin only
func AdminOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		roleID := c.Get("role_id")
		if roleID != uint(1) && roleID != 1 {
			return c.JSON(http.StatusForbidden, echo.Map{"error": "admin only access"})
		}
		return next(c)
	}
}
