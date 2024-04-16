package delivery

import (
	"koi-backend-web-go/domain"
	"koi-backend-web-go/middleware"

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
	private.Get("/get-event-by-mahasiswa", middleware.ValidateTokenMahasiswa, handler.GetEventByMahasiswaID)

	_ = api.Group("/public")

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
