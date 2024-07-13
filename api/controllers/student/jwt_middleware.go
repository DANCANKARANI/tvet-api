package student

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/middleware"
	"github.com/pentabyte/tvet/api/api/utilities"
)

func JWTMiddleware(c *fiber.Ctx) error {
// Check for token in cookies first
tokenString := c.Cookies("Authorization")
 str := c.Get("Authorization")
 fmt.Println(str)
	if tokenString == ""{
		log.Println("missing jwt")
		return utilities.ShowError(c,"unauthorized",fiber.StatusUnauthorized)
	}
	//validate the token
	claims,err :=middleware.ValidateToken(tokenString)
	if err != nil{
		log.Println(err.Error())
		utilities.ShowError(c,"unauthorized",fiber.StatusUnauthorized)
	}
	//store the userID 
	c.Locals("user_id",claims.UserID)
	return c.Next()
}