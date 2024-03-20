package delivery

import (
	"koi-backend-web-go/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pandeptwidyaop/golog"
)

type DetailKegiatanHandler struct {
	DetailKegiatanUC domain.DetailKegiatanUseCase
}

func NewDetailKegiatanHandler(c *fiber.App, das domain.DetailKegiatanUseCase) {
	handler := &DetailKegiatanHandler{
		DetailKegiatanUC: das,
	}
	api := c.Group("/detail-kegiatan")

	_ = api.Group("/private")

	public := api.Group("/public")
	public.Get("/get-by-id/:id", handler.GetDetailKegiatanByID)
}

func (t *DetailKegiatanHandler) GetDetailKegiatanByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": "failed convert to int",
			"error":   err,
		})
	}
	res, er := t.DetailKegiatanUC.GetDetailKegiatanByID(c.Context(), uint(id))
	if er != nil {
		golog.Slack.ErrorWithData("error get detail kegiatan", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully get detail kegiatan",
	})
}
