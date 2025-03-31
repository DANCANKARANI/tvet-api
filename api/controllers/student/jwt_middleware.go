package student

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/middleware"
	"github.com/pentabyte/tvet/api/api/utilities"
)


func JWTMiddleware(c *fiber.Ctx) error {
    // Try extracting token from cookies first
    tokenString := c.Cookies("Authorization")

    // If token not in cookies, check Authorization header
    authHeader := c.Get("Authorization")
    if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
        tokenString = strings.TrimPrefix(authHeader, "Bearer ")
    }

    // If no token found, return 401 Unauthorized
    if tokenString == "" {
        log.Println("Missing JWT")
        return utilities.ShowError(c, "Unauthorized", fiber.StatusUnauthorized)
    }
 log.Print(tokenString)
    // Validate the token
    claims, err := middleware.ValidateToken(tokenString)
    if err != nil {
        log.Println("JWT validation failed:", err.Error())
        return utilities.ShowError(c, "Unauthorized", fiber.StatusUnauthorized)
    }

    // Store the user ID in request context
    c.Locals("user_id", claims.UserID)

    // Continue processing request
    return c.Next()
}