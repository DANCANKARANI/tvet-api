package sponsors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pentabyte/tvet/api/api/controllers/sponsor"
	//"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetSponsorsRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000, http://127.0.0.1",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/sponsor")
	sponsorGroup := auth.Group("/")
	sponsorGroup.Get("/",sponsor.GetSponsorHandler)
	sponsorGroup.Post("/",sponsor.AddSponsorHandler )
	sponsorGroup.Patch("/:id",sponsor.UpdateSponsorHandler)
	sponsorGroup.Delete("/:id",sponsor.DeleteSponsorHandler)
}