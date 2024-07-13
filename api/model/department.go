package model

import (
	"errors"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

)

/*
Add all the departments
@params c *fiber.Ctx
*/
func AddDepartment(c *fiber.Ctx)(*Department,error){
	id := uuid.New()
	body := Department{}
	response :=Department{}
	if err :=c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add departments")
	}
	body.ID=id
	err := db.Create(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add departments")
	}
	return &response,nil
}
/*
updates the departments
@params department_id
*/
func UpdateDepartment(c *fiber.Ctx,department_id uuid.UUID)(*Department,error){
	body:=Department{}
	response := Department{}
	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update department")
	}
	err:=db.Model(&Department{}).Where("id = ?",department_id).Updates(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update department")
	}
	return &response,nil
}
/*
delete departments
@params department_id
*/
func DeleteDepartment(job_id uuid.UUID)error{
	err:= db.First(&Department{},"id = ?",job_id).Delete(&Department{}).Error
	if err != nil{
		log.Println(err.Error())
		return errors.New("failed to delete departments")
	}
	return nil
}
//gets all the jobs
func GetDepartment()(*[]Department,error){
	response:=[]Department{}
	err:=db.Model(&Department{}).Scan(&response).Error
	if err != nil{
		log.Println(err.Error())
		return nil,errors.New("failed to get departments")
	}
	return &response,nil
}