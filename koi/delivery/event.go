package delivery

import (
	"fmt"
	"koi-backend-web-go/domain"
	"koi-backend-web-go/middleware"
	"strconv"

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
	private.Post("/create-event", middleware.ValidateTokenOrmawa, handler.CreateEvent)
	private.Put("/update-event", middleware.Validate, handler.UpdateEvent)
	private.Get("/get-event-ormawa", middleware.ValidateTokenOrmawa, handler.GetAllEventsIdOrmawaSide)
	private.Get("/get-event-by-id-and-ormawa/:id", middleware.Validate, handler.GetEventByIDAndOrmawaID)
	private.Delete("/delete-event/:id", middleware.Validate, handler.DeleteEvent)

	private.Put("/update-status-event/:id", middleware.Validate, handler.UpdateStatusEvent)

	public := api.Group("/public")
	public.Get("/get-all-events", handler.GetAllEvents)
	public.Get("/get-event-by-ormawa/:id", handler.GetAllEventsIdOrmawa)
	public.Get("/get-event-by-id/:id", handler.GetEventByID)
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
	local := c.Locals("role")
	fmt.Println("ini local brok: ", local)
	res, er := t.EventUC.CreateEvent(c.Context(), req)
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
		"local":   local,
		"message": "Successfully create event",
	})
}

func (t *EventHandler) UpdateEvent(c *fiber.Ctx) error {
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
	local := c.Locals("role")
	fmt.Println("ini local brok: ", local)
	er := t.EventUC.UpdateEvent(c.Context(), req)
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
		"local":   local,
		"message": "Successfully create event",
	})
}

func (t *EventHandler) GetAllEvents(c *fiber.Ctx) error {
	res, er := t.EventUC.GetAllEvents(c.Context())
	if er != nil {
		golog.Slack.ErrorWithData("error get event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully get event",
	})
}

func (t *EventHandler) GetAllEventsIdOrmawaSide(c *fiber.Ctx) error {
	id := middleware.UserID(c)
	res, er := t.EventUC.GetEventByOrmawaID(c.Context(), uint(id))
	if er != nil {
		golog.Slack.ErrorWithData("error get event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully get event",
	})
}

func (t *EventHandler) GetAllEventsIdOrmawa(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": "failed convert to int",
			"error":   err,
		})
	}
	res, er := t.EventUC.GetEventByOrmawaID(c.Context(), uint(id))
	if er != nil {
		golog.Slack.ErrorWithData("error get event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully get event",
	})
}

func (t *EventHandler) GetEventByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": "failed convert to int",
			"error":   err,
		})
	}
	res, er := t.EventUC.GetEventByID(c.Context(), uint(id))
	if er != nil {
		golog.Slack.ErrorWithData("error get event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully get event",
	})
}

func (t *EventHandler) GetEventByIDAndOrmawaID(c *fiber.Ctx) error {
	eventId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": "failed convert to int",
			"error":   err,
		})
	}
	ormawaID := middleware.UserID(c)
	res, er := t.EventUC.GetEventByIDAndOrmawaID(c.Context(), uint(ormawaID), uint(eventId))
	if er != nil {
		golog.Slack.ErrorWithData("error get event", c.Body(), er)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": er,
			"error":   er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"message": "Successfully get event",
	})
}

func (t *EventHandler) DeleteEvent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": "failed convert to int",
			"error":   err,
		})
	}
	er := t.EventUC.DeleteEvent(c.Context(), uint(id))
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

func (t *EventHandler) UpdateStatusEvent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": "failed convert to int",
			"error":   err,
		})
	}
	req := new(domain.ChangeStatusRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	er := t.EventUC.UpdateStatusEvent(c.Context(), uint(id), req.ItsOpen)
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
		"message": "Successfully update status event",
	})
}
