
package model

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
gets user using user ID
*/
type ResponseStudent struct{
	ID uuid.UUID 		`json:"id"`
	FullName string 	`json:"full_name"`
	PhoneNumber string 	`json:"phone_number"`
	Email string 		`json:"email"`
}

func GetOneUSer(c *fiber.Ctx)(*ResponseStudent,error){
	id,err:=GetAuthUserID(c)
	if err != nil {
		return nil,errors.New("failed to get user's id:"+err.Error())
	}
	user := ResponseStudent{}
	err = db.Preload("Image").First(&Student{},"id = ?",id).Scan(&user).Error
	if err != nil {
		return nil,errors.New("failed to get user details:"+err.Error())
	}
	return &user,nil
}
//gets all the users
func GetAllUsersDetails(c *fiber.Ctx)(*[]ResponseStudent,error){
	response:=[]ResponseStudent{}
	err := db.Model(&Student{}).Scan(&response).Error
	if err != nil {
		return nil,errors.New("failed to get users:"+err.Error())
	}
	return &response,nil
}
//updates the user by id
func UpdateUser(c *fiber.Ctx)(*ResponseStudent,error){
	id,err:=GetAuthUserID(c)
	if err != nil {
		return nil,errors.New("failed to get user's id:"+err.Error())
	}
	body := Student{}
	if err = c.BodyParser(&body);err != nil {
		return nil,errors.New("failed to parse:"+err.Error())
	}
	response := ResponseStudent{}
	err = db.First(&Student{},"id = ?",id).Updates(&body).Scan(&response).Error
	if err != nil {
		return nil,errors.New("error in updating the user:"+err.Error())
	}
	return &response,nil
}

