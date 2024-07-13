package sponsors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pentabyte/tvet/api/api/controllers/sponsor"
	"github.com/pentabyte/tvet/api/api/controllers/student"
)

func SetSponsorsRoutes(app *fiber.App) {
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/sponsor")
	sponsorGroup := auth.Group("/",student.JWTMiddleware)
	sponsorGroup.Get("/",sponsor.GetSponsorHandler)
	sponsorGroup.Post("/",sponsor.AddSponsorHandler )
	sponsorGroup.Patch("/:id",sponsor.UpdateSponsorHandler)
	sponsorGroup.Delete("/:id",sponsor.DeleteSponsorHandler)
}