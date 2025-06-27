package http

import (
	"github.com/YurcheuskiRadzivon/test-to-do-api-gateway/config"
	"github.com/gofiber/fiber/v2"
)

type APIController struct {
	app *fiber.App
	cfg *config.Config
}
