package model

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResponseJobs struct {
	ID              uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key"`
	Title           string         `json:"title" gorm:"size:255"`
	Role            string         `json:"role" gorm:"type:varchar(255)"`
	ApplicationLink string         `json:"application_link" gorm:"type:varchar(1024)"`
	CreatedAt       time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
/*
Add all the jobs
@params c *fiber.Ctx
*/
func AddJobs(c *fiber.Ctx)(*ResponseJobs,error){
	id := uuid.New()
	body := Job{}
	response :=ResponseJobs{}
	if err :=c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add jobs")
	}
	body.ID=id
	err := db.Create(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add jobs")
	}
	return &response,nil
}
/*
updates the jobs
@params job id
*/
func UpdateJob(c *fiber.Ctx,job_id uuid.UUID)(*ResponseJobs,error){
	body:=Job{}
	response := ResponseJobs{}
	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update jobs")
	}
	err:=db.Model(&Job{}).Where("id = ?",job_id).Updates(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update jobs")
	}
	return &response,nil
}
/*
delete jobs
@params job_id
*/
func DeleteJob(job_id uuid.UUID)error{
	err:= db.First(&Job{},"id = ?",job_id).Delete(&Job{}).Error
	if err != nil{
		log.Println(err.Error())
		return errors.New("failed to delete job")
	}
	return nil
}
//gets all the jobs
func GetJobs()(*[]ResponseJobs,error){
	job:=[]ResponseJobs{}
	err:=db.Model(&Job{}).Scan(&job).Error
	if err != nil{
		log.Println(err.Error())
		return nil,errors.New("failed to get jobs")
	}
	return &job,nil
}