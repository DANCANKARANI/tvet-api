package jobs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pentabyte/tvet/api/api/controllers/job"
	//"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetJobsRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000, http://127.0.0.1",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/job")
	jobGroup := auth.Group("/")
	jobGroup.Get("/",job.GetJobHandler)
	jobGroup.Post("/",job.AddJobHandler )
	jobGroup.Patch("/:id",job.UpdateJobHandler)
	jobGroup.Delete("/:id",job.DeleteJobHandler)
}