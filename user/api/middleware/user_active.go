package middleware

// func ActiveOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		roleID := c.Get("role_id")
// 		userID := c.Get("user_id")

// 		// ambil user dari DB kalau perlu, atau inject saat login
// 		user, err := repositoryInstance.FindByID(userID.(uint))
// 		if err != nil || !user.IsActive {
// 			return c.JSON(http.StatusForbidden, echo.Map{"error": "your account is inactive"})
// 		}

// 		return next(c)
// 	}
// }
