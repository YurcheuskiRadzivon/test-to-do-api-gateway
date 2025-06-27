package http

import (
	"github.com/YurcheuskiRadzivon/test-to-do-api-gateway/config"
	"github.com/gofiber/fiber/v2"
)

func NewRoute(
	app *fiber.App,
	cfg *config.Config,
) {
	APIC := &APIController{
		app: app,
		cfg: cfg,
	}

	_ = APIC

	authGroup := app.Group("/auth")
	{
	}

	_ = authGroup

	adminGroup := app.Group("/admin")
	{
	}

	_ = adminGroup

	userGroup := app.Group("/account")
	{
	}

	_ = userGroup

	noteGroup := app.Group("/note")
	{
	}

	_ = noteGroup

	fileGroup := app.Group("/file")
	{
	}

	_ = fileGroup

}
