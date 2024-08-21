package student

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/utilities"
	"github.com/pentabyte/tvet/api/api/model"
)

func GetOneStudentHandler(c *fiber.Ctx) error {
	user,err := model.GetOneUSer(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"user retrieved successfully",fiber.StatusOK,user)
}
func UpdateStudentHandler(c *fiber.Ctx)error{
	response,err := model.UpdateUser(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"user updated successfully",fiber.StatusOK,response)
}

func GetAllStudent(c *fiber.Ctx)error{
	response, err := model.GetAllUsersDetails(c)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully retrieved all students",fiber.StatusOK, response)
}