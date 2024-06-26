package delivery

import (
	"koi-backend-web-go/domain"
	"koi-backend-web-go/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pandeptwidyaop/golog"
)

type PembayaranHandler struct {
	PembayaranUC domain.PembayaranUseCase
}

func NewPembayaranHandler(c *fiber.App, das domain.PembayaranUseCase) {
	handler := &PembayaranHandler{
		PembayaranUC: das,
	}
	api := c.Group("/pembayaran")

	private := api.Group("/private")
	private.Post("/create", middleware.ValidateTokenMahasiswa, handler.CreatePemabayaran)
	private.Put("/update/:id", middleware.ValidateTokenOrmawa, handler.CreatePemabayaran)
	private.Put("/update-status/:id", middleware.ValidateTokenOrmawa, handler.UpdatePembayaran)
	// delete
	private.Delete("/delete/:id", middleware.ValidateTokenOrmawa, handler.DeletePembayaran)

	private.Get("/get-event-by-mahasiswa", middleware.ValidateTokenMahasiswa, handler.GetEventByMahasiswaID)

	public := api.Group("/public")
	public.Get("/get-event", handler.GetEvents)
}

func (t *PembayaranHandler) CreatePemabayaran(c *fiber.Ctx) error {
	req := new(domain.Pembayaran)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	// fmt.Println("ini isian variable c: ", c)
	id := middleware.UserID(c)
	req.MahasiswaID = uint(id)
	// fmt.Println("ini id ormawa: ", req)
	res, er := t.PembayaranUC.CreatePembayaran(c.Context(), req)
	if er != nil {
		golog.Slack.ErrorWithData("error create event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully create event",
	})
}

func (t *PembayaranHandler) GetEventByMahasiswaID(c *fiber.Ctx) error {
	// fmt.Println("ini isian variable c: ", c)
	id := middleware.UserID(c)
	// fmt.Println("ini id ormawa: ", req)
	res, er := t.PembayaranUC.GetEventByMahasiswaID(c.Context(), uint(id))
	if er != nil {
		golog.Slack.ErrorWithData("error get event by mahasiswa id", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully get event",
	})
}

func (t *PembayaranHandler) GetEvents(c *fiber.Ctx) error {
	res, er := t.PembayaranUC.GetEvents(c.Context())
	if er != nil {
		golog.Slack.ErrorWithData("error get event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully get event",
	})
}

func (t *PembayaranHandler) UpdatePembayaran(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse id",
			"error":   err,
		})
	}
	req := new(domain.Pembayaran)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	req.ID = uint(id)
	res, er := t.PembayaranUC.UpdatePembayaran(c.Context(), req)
	if er != nil {
		golog.Slack.ErrorWithData("error update event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully update event",
	})
}

func (t *PembayaranHandler) UpdateStatusPembayaran(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse id",
			"error":   err,
		})
	}
	req := new(domain.Pembayaran)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	req.ID = uint(id)
	res, er := t.PembayaranUC.UpdateStatusPembayaran(c.Context(), req)
	if er != nil {
		golog.Slack.ErrorWithData("error update status event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully update status event",
	})
}

func (t *PembayaranHandler) DeletePembayaran(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse id",
			"error":   err,
		})
	}
	er := t.PembayaranUC.DeletePembayaran(c.Context(), uint(id))
	if er != nil {
		golog.Slack.ErrorWithData("error delete event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Successfully delete event",
	})
}
