package courses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/controllers/course"
	"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetCoursesRoutes(app *fiber.App) {
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/course")
	courseGroup := auth.Group("/",student.JWTMiddleware)
	courseGroup.Get("/",course.GetCourseHandler)
	courseGroup.Post("/",course.AddCourseHandler )
	courseGroup.Patch("/:id",course.UpdateCourseHandler)
	courseGroup.Delete("/:id",course.DeleteCourseHandler)
}