package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware checks if the user's JWT is valid and if their role matches one of the allowed roles
func AuthMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get token from Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
			}

			// Extract token from "Bearer <token>"
			tokenString := strings.Split(authHeader, " ")[1]
			log.Println("Received Token:", tokenString) // Debug log

			// Parse JWT token and get the claims
			claims := &jwt.MapClaims{}
			accessSecret := os.Getenv("ACCESS_SECRET") // Get the secret key from environment variables

			// Parse and validate the token using the secret key
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(accessSecret), nil
			})

			if err != nil || !token.Valid {
				log.Println("Error Parsing Token:", err) // Debug log
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}

			// Print claims for debugging
			log.Println("Token Claims:", claims) // Debug log

			// Extract the role from the claims
			role, ok := (*claims)["role"].(string)
			if !ok {
				log.Println("Error: role not found in token") // Debug log
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "role not found"})
			}

			// Check if the user's role matches one of the allowed roles
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					return next(c)
				}
			}

			// If the role is not authorized
			log.Println("Error: insufficient privileges") // Debug log
			return c.JSON(http.StatusForbidden, map[string]string{"error": "insufficient privileges"})
		}
	}
}
