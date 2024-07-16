
package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/routes/courses"
	"github.com/pentabyte/tvet/api/api/routes/departments"
	"github.com/pentabyte/tvet/api/api/routes/jobs"
	"github.com/pentabyte/tvet/api/api/routes/sponsors"
	"github.com/pentabyte/tvet/api/api/routes/students"
)

func CreateEndpoint() {
	app := fiber.New()
	
	
	students.SetStudentRoutes(app)
	jobs.SetJobsRoutes(app)
	sponsors.SetSponsorsRoutes(app)
	courses.SetCoursesRoutes(app)
	departments.SetDepartmentsRoutes(app)
	app.Listen(":3000")
}