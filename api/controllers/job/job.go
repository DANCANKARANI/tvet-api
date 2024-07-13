package job

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pentabyte/tvet/api/api/model"
	"github.com/pentabyte/tvet/api/api/utilities"
)

//add job handler
func AddJobHandler(c *fiber.Ctx) error {
	response,err:= model.AddJobs(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully added jobs",fiber.StatusOK,response)
}
//update Job handler
func UpdateJobHandler(c *fiber.Ctx)error{
	job_id,_:=uuid.Parse(c.Params("id"))
	response,err:=model.UpdateJob(c,job_id)
	if err!=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated jobs",fiber.StatusOK,response)
}
//get jobs handler
func GetJobHandler(c *fiber.Ctx)error{
	response,err:=model.GetJobs()
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfuly retrieved jobs",fiber.StatusOK,response)
}
//delete job handler
func DeleteJobHandler(c *fiber.Ctx)error{
	job_id,_:=uuid.Parse(c.Params("id"))
	err:=model.DeleteJob(job_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully deleted job",fiber.StatusOK)
}