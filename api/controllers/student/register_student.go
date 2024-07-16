package student

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pentabyte/tvet/api/api/database"
	"github.com/pentabyte/tvet/api/api/model"
	"github.com/pentabyte/tvet/api/api/utilities"
)


var db =database.ConnectDB()
var country_code = "KE"

func CreateStudentAccount(c *fiber.Ctx) error {
	db.AutoMigrate(&model.Student{})
	//generating new id
	id := uuid.New()
	student:=model.Student{}
	if err :=c.BodyParser(&student); err != nil {
		log.Fatal(err.Error())
		return utilities.ShowError(c,"failed to create account", fiber.StatusInternalServerError)
	}

	//validate email address
	if ! utilities.ValidateEmail(student.Email){
		errStr:="inavalid email address "+student.Email
		return utilities.ShowError(c,errStr, fiber.StatusNotAcceptable)
	}
	//Check if user exist
	userExist,_,err:= model.UserExist(c,student.PhoneNumber)
	if err != nil{
		return utilities.ShowError(c,err.Error(),fiber.StatusInternalServerError)
	}
	if userExist{
		errStr:="User with this phone number "+student.PhoneNumber+" already exists"
		return utilities.ShowError(c,errStr,fiber.StatusConflict)
	}
	//validate phone number
	_,err = utilities.ValidatePhoneNumber(student.PhoneNumber,country_code)
	if err !=nil{
		log.Fatal(err.Error())
		return utilities.ShowError(c,err.Error(),fiber.StatusAccepted)
	}

	
	//hash password
	hashed_password, _:= utilities.HashPassword(student.Password)

	userModel := model.Student{ID: id,FullName: student.FullName,Email: student.Email,PhoneNumber: student.PhoneNumber,Password: hashed_password}
	//create user
	//userModel.CodeExpirationTime=time.Now()
	err = db.Create(&userModel).Error
	if err!= nil {
		log.Fatal(err.Error())
		return utilities.ShowError(c, "failed to add data to the database",fiber.StatusInternalServerError)
	}
	return utilities.ShowMessage(c,"account created successfully",fiber.StatusOK)
}