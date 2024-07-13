package model

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResponseSponsor struct {
	ID              uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key"`
	Name            string         `json:"name" gorm:"size:255"`
	Description     string         `json:"descrption" gorm:"type:text"`
	ApplicationLink string         `json:"application_link" gorm:"type:varchar(1024)"`
	CreatedAt       time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
/*
Add all the sponsors
@params c *fiber.Ctx
*/
func AddSponsor(c *fiber.Ctx)(*ResponseSponsor,error){
	id := uuid.New()
	body := Sponsor{}
	response :=ResponseSponsor{}
	if err :=c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add payament")
	}
	body.ID=id
	err := db.Create(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed sponsor")
	}
	return &response,nil
}
/*
updates the sponsors
@params sponsor id
*/
func UpdateSponsor(c *fiber.Ctx,sponsor_id uuid.UUID)(*ResponseSponsor,error){
	body:=Sponsor{}
	response := ResponseSponsor{}
	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update sponsors")
	}
	err:=db.First(&Sponsor{},"id = ?",sponsor_id).Updates(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update sponsors")
	}
	return &response,nil
}
/*
delete sponsors
@params sponsor_id
*/
func DeleteSponsor(sponsor_id uuid.UUID)error{
	err:= db.First(&Sponsor{},"id = ?",sponsor_id).Delete(&Sponsor{}).Error
	if err != nil{
		log.Println(err.Error())
		return errors.New("failed to delete sponsor")
	}
	return nil
}
//gets all the sponsors
func GetSponsors()(*[]ResponseSponsor,error){
	job:=[]ResponseSponsor{}
	err:=db.Model(&Sponsor{}).Scan(&job).Error
	if err != nil{
		log.Println(err.Error())
		return nil,errors.New("failed to get sponsors")
	}
	return &job,nil
}