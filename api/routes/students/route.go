package students

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetStudentRoutes(app *fiber.App) {
	
	auth := app.Group("/api/v1/student")
	auth.Post("/",student.CreateStudentAccount)
	auth.Post("/login",student.Login)
	//protected routes
	userGroup := auth.Group("/",student.JWTMiddleware)
	userGroup.Get("/",student.GetOneStudentHandler)
}