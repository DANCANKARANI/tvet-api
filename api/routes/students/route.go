package students

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetStudentRoutes(app *fiber.App) {
	
	auth := app.Group("/api/v1/student")
	auth.Get("/all",student.GetAllStudent)
	auth.Post("/register",student.CreateStudentAccount)
	auth.Post("/login",student.Login)
	authG:=auth.Group("/",student.JWTMiddleware)
	authG.Get("/",student.GetOneStudentHandler)
	}