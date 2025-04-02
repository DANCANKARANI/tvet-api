package model

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/utilities"
)
type ContatData struct{
	Name string 	`json:"name"`
	Email	string	`json:"email"`	
	Message string	`json:"message"`
}

func ContactUs(c *fiber.Ctx) error {
	body := ContatData{}
	if err:= c.BodyParser(&body); err != nil{
		log.Println(err.Error())
		return utilities.ShowError(c,"failed to parse json data",fiber.StatusInternalServerError)
	}
	err := utilities.SendMail(body.Name,body.Email,body.Message)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"message sent successfully",fiber.StatusOK,body)
}

