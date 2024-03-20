package delivery

import (
	"fmt"
	"koi-backend-web-go/domain"
	"koi-backend-web-go/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/pandeptwidyaop/golog"
)

type EventHandler struct {
	EventUC domain.EventUseCase
}

func NewEventHandler(c *fiber.App, das domain.EventUseCase) {
	handler := &EventHandler{
		EventUC: das,
	}
	api := c.Group("/event")
	private := api.Group("/private")
	private.Post("/create-event", middleware.ValidateToken, handler.CreateEvent)
	_ = api.Group("/public")

}

func (t *EventHandler) CreateEvent(c *fiber.Ctx) error {
	req := new(domain.CreateEvent)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	fmt.Println("ini isian variable c: ", c)
	id := middleware.UserID(c)
	req.OrmawaID = uint(id)
	fmt.Println("ini id ormawa: ", req)
	res, er := t.EventUC.CreateEvent(c.Context(), req)
	if er != nil {
		golog.Slack.ErrorWithData("error create event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully create event",
	})
}
