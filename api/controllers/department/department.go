package department

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pentabyte/tvet/api/api/model"
	"github.com/pentabyte/tvet/api/api/utilities"
)

//add department handler
func AddDepartmentHandler(c *fiber.Ctx) error {
	response,err:= model.AddDepartment(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully added department",fiber.StatusOK,response)
}
//update department handler
func UpdateDepartmentHandler(c *fiber.Ctx)error{
	department_id,_:=uuid.Parse(c.Params("id"))
	response,err:=model.UpdateDepartment(c,department_id)
	if err!=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated department",fiber.StatusOK,response)
}
//get jobs handler
func GetDepartmentHandler(c *fiber.Ctx)error{
	response,err:=model.GetDepartment()
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfuly retrieved departments",fiber.StatusOK,response)
}
//delete job handler
func DeleteDepartmentHandler(c *fiber.Ctx)error{
	department_id,_:=uuid.Parse(c.Params("id"))
	err:=model.DeleteDepartment(department_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully deleted department",fiber.StatusOK)
}