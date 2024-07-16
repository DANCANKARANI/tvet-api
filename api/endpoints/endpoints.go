package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pentabyte/tvet/api/api/routes/courses"
	"github.com/pentabyte/tvet/api/api/routes/departments"
	"github.com/pentabyte/tvet/api/api/routes/jobs"
	"github.com/pentabyte/tvet/api/api/routes/sponsors"
	"github.com/pentabyte/tvet/api/api/routes/students"
)

func CreateEndpoint() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000, http://127.0.0.1,https://idristvet.vercel.app,*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
	students.SetStudentRoutes(app)
	jobs.SetJobsRoutes(app)
	sponsors.SetSponsorsRoutes(app)
	courses.SetCoursesRoutes(app)
	departments.SetDepartmentsRoutes(app)
	app.Listen(":3000")
}