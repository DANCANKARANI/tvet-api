package courses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pentabyte/tvet/api/api/controllers/course"
	"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetCoursesRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000, http://127.0.0.1",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/course")
	courseGroup := auth.Group("/",student.JWTMiddleware)
	courseGroup.Get("/",course.GetCourseHandler)
	courseGroup.Post("/",course.AddCourseHandler )
	courseGroup.Patch("/:id",course.UpdateCourseHandler)
	courseGroup.Delete("/:id",course.DeleteCourseHandler)
}