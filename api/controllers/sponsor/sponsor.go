package sponsor

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pentabyte/tvet/api/api/model"
	"github.com/pentabyte/tvet/api/api/utilities"
)

//add job handler
func AddSponsorHandler(c *fiber.Ctx) error {
	response,err:= model.AddSponsor(c)
	if err != nil {
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully added sponsors",fiber.StatusOK,response)
}
//update Job handler
func UpdateSponsorHandler(c *fiber.Ctx)error{
	sponsor_id,_:=uuid.Parse(c.Params("id"))
	response,err:=model.UpdateSponsor(c,sponsor_id)
	if err!=nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfully updated sponsors",fiber.StatusOK,response)
}
//get jobs handler
func GetSponsorHandler(c *fiber.Ctx)error{
	response,err:=model.GetSponsors()
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowSuccess(c,"successfuly retrieved sponsors",fiber.StatusOK,response)
}
//delete job handler
func DeleteSponsorHandler(c *fiber.Ctx)error{
	sponsor_id,_:=uuid.Parse(c.Params("id"))
	err:=model.DeleteSponsor(sponsor_id)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"successfully deleted sponsor",fiber.StatusOK)
}