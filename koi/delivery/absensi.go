package delivery

import (
	"fmt"
	"koi-backend-web-go/domain"
	"koi-backend-web-go/middleware"
	"strconv"

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
	// get all absensi
	public.Get("/get-all", handler.GetAllAbsensiHandler)
	public.Get("/get-absent/:id", handler.GetAbsensiByEventID)
	public.Put("/update/status/:id", middleware.Validate, handler.UpdateStatusHandler)
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

func (t *AbsensiHandler) GetAbsensiByEventID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": "failed convert to int",
			"error":   err,
		})
	}

	res, err := t.AbsensiUC.GetAbsensiByEventID(c.Context(), uint(id))
	if err != nil {
		golog.Slack.ErrorWithData("error create user", c.Body(), err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully create user",
	})
}

func (t *AbsensiHandler) UpdateStatusHandler(c *fiber.Ctx) error {
	req := new(domain.Absensi)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	idEvent, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": "failed convert to int",
			"error":   err,
		})
	}
	fmt.Println(req.UserId)
	er := t.AbsensiUC.UpdateStatus(c.Context(), uint(idEvent), req.UserId, req.Status)
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
		"message": "Successfully create user",
	})
}

func (t *AbsensiHandler) GetAllAbsensiHandler(c *fiber.Ctx) error {
	res, err := t.AbsensiUC.GetAllAbsensi(c.Context())
	if err != nil {
		golog.Slack.ErrorWithData("error create user", c.Body(), err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully create user",
	})
}
