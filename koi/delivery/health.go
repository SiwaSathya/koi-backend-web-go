package delivery

import (
	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler(c *fiber.App) {
	handler := &HealthCheckHandler{}

	c.Get("/", handler.HealthCheck)
}

func (h *HealthCheckHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "OK",
	})
}
