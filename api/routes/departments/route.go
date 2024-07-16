package departments

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pentabyte/tvet/api/api/controllers/department"
	"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetDepartmentsRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000, http://127.0.0.1",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/department")
	departmentGroup := auth.Group("/",student.JWTMiddleware)
	departmentGroup.Get("/",department.GetDepartmentHandler)
	departmentGroup.Post("/",department.AddDepartmentHandler )
	departmentGroup.Patch("/:id",department.UpdateDepartmentHandler)
	departmentGroup.Delete("/:id",department.DeleteDepartmentHandler)
}