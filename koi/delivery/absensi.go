package delivery

import (
	"koi-backend-web-go/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/pandeptwidyaop/golog"
)

type AbsensiHandler struct {
	AbsensiUC domain.AbsensiUseCase
}

func NewAbsensiHandler(c *fiber.App, das domain.AbsensiUseCase) {
	handler := &AbsensiHandler{
		AbsensiUC: das,
	}
	api := c.Group("/absensi")

	_ = api.Group("/private")
	public := api.Group("/public")
	public.Post("/create", handler.AbsensiHandler)
}

func (t *AbsensiHandler) AbsensiHandler(c *fiber.Ctx) error {
	req := new(domain.Absensi)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	res, er := t.AbsensiUC.CreateAbsensi(c.Context(), req)
	if er != nil {
		golog.Slack.ErrorWithData("error create user", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully create user",
	})
}
