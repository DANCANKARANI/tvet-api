package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pentabyte/tvet/api/api/database"
	"gorm.io/gorm"
)
var db =database.ConnectDB()
/*
finds user using phone number only
@params phone_number
*/

func UserExist(c *fiber.Ctx, phoneNumber string) (bool, *Student, error) {
    var existingUser Student

    // Detailed logging
    log.Printf("Checking for user with phone number: %s", phoneNumber)

    err := db.Where("phone_number = ?", phoneNumber).First(&existingUser).Error
    if err != nil {
        // Log the detailed error
        log.Printf("Error finding user with phone number %s: %v", phoneNumber, err.Error())

        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false, nil, nil
        }

        return false, nil, fmt.Errorf("database error: %v", err.Error())
    }
	log.Printf("User found: %+v", existingUser)
    return true, &existingUser, nil
}
//gets user id
func GetAuthUserID(c *fiber.Ctx)(uuid.UUID,error){
	user_id :=c.Locals("user_id")
	id,ok := user_id.(*uuid.UUID)
	if !ok{
		return uuid.Nil,errors.New("unauthorized")
	}
	user_id=*id
	return user_id.(uuid.UUID),nil
}