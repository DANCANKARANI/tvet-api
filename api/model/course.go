package model

import (
	"errors"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
Add all the Course
@params c *fiber.Ctx
*/
var body Course
func AddCourse(c *fiber.Ctx)(*Course,error){
	id := uuid.New()
	
	response :=Course{}
	if err :=c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add course")
	}
	body.ID=id
	err := db.Create(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to add course")
	}
	return &response,nil
}
/*
updates the course
@params course id
*/
func UpdateCourse(c *fiber.Ctx,course_id uuid.UUID)(*Course,error){
	response := Course{}
	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update course")
	}
	err:=db.Model(&Course{}).Where("id = ?",course_id).Updates(&body).Scan(&response).Error
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New("failed to update course")
	}
	return &response,nil
}
/*
delete courses
@params course_id
*/
func DeleteCourse(course_id uuid.UUID)error{
	err:= db.First(&Course{},"id = ?",course_id).Delete(&Course{}).Error
	if err != nil{
		log.Println(err.Error())
		return errors.New("failed to delete course")
	}
	return nil
}
//gets all the jobs
func GetCourses()(*[]Course,error){
	course:=[]Course{}
	err:=db.Model(&Course{}).Scan(&course).Error
	if err != nil{
		log.Println(err.Error())
		return nil,errors.New("failed to get course")
	}
	return &course,nil
}