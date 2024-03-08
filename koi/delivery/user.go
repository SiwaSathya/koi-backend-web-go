package delivery

import (
	"koi-backend-web-go/domain"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/pandeptwidyaop/golog"
)

type UserHandler struct {
	UserUC domain.UserUseCase
}

func NewUserHandler(c *fiber.App, das domain.UserUseCase) {
	handler := &UserHandler{
		UserUC: das,
	}
	api := c.Group("/user")
	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)

}

func (t *UserHandler) Register(c *fiber.Ctx) error {
	req := new(domain.CreateUser)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	res, er := t.UserUC.CreateUser(c.Context(), req)
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

func (t *UserHandler) Login(c *fiber.Ctx) error {
	req := new(domain.LoginPayload)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	valRes, er := govalidator.ValidateStruct(req)
	if !valRes {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   er.Error(),
		})
	}
	res, token, er := t.UserUC.LoginUser(c.Context(), req)
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
		"token":   token,
		"message": "Successfully create user",
	})
}
