package students

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetStudentRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000, http://127.0.0.1",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
	auth := app.Group("/api/v1/student")
	auth.Post("/register",student.CreateStudentAccount)
	auth.Post("/login",student.Login)
	auth.Get("/all",student.GetAllStudent)
	//protected routes
	userGroup := auth.Group("/",student.JWTMiddleware)
	userGroup.Get("/one",student.GetOneStudentHandler)
}