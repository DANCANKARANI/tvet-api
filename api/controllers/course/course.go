package course

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pentabyte/tvet/api/api/model"
	"github.com/pentabyte/tvet/api/api/utilities"
)

//add course handler
func AddCourseHandler(c *fiber.Ctx) error {
	response,err:= model.AddCourse(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully added course",fiber.StatusOK,response)
}
//update course handler
func UpdateCourseHandler(c *fiber.Ctx)error{
	course_id,_:=uuid.Parse(c.Params("id"))
	response,err:=model.UpdateCourse(c,course_id)
	if err!=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated course",fiber.StatusOK,response)
}
//get course handler
func GetCourseHandler(c *fiber.Ctx)error{
	response,err:=model.GetCourses()
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfuly retrieved courses",fiber.StatusOK,response)
}
//delete course handler
func DeleteCourseHandler(c *fiber.Ctx)error{
	course_id,_:=uuid.Parse(c.Params("id"))
	err:=model.DeleteCourse(course_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully deleted course",fiber.StatusOK)
}