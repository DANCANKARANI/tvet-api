package jobs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/controllers/job"
	"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetJobsRoutes(app *fiber.App) {
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/job")
	jobGroup := auth.Group("/",student.JWTMiddleware)
	jobGroup.Get("/",job.GetJobHandler)
	jobGroup.Post("/",job.AddJobHandler )
	jobGroup.Patch("/:id",job.UpdateJobHandler)
	jobGroup.Delete("/:id",job.DeleteJobHandler)
}