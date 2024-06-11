package delivery

import (
	"koi-backend-web-go/domain"
	"koi-backend-web-go/middleware"

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
	private := api.Group("/private")
	private.Get("/profile", middleware.Validate, handler.GetProfie)
	private.Get("/dashboard-kemahasiswaan", middleware.ValidateTokenKemahasiswaan, handler.GetDashboardKemahasiswaan)
	private.Post("/reset-password", middleware.Validate, handler.ResetPassword)
	private.Put("/edit-profile/:role", middleware.Validate, handler.EditProfile)
	public := api.Group("/public")
	public.Post("/register", handler.Register)
	public.Post("/login", handler.Login)

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

func (t *UserHandler) EditProfile(c *fiber.Ctx) error {
	req := new(domain.CreateUser)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	role := c.Params("role")
	req.Role = role
	er := t.UserUC.UpdateProfile(c.Context(), req)
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

func (t *UserHandler) ResetPassword(c *fiber.Ctx) error {
	req := new(domain.ResetPassword)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	er := t.UserUC.ResetPassword(c.Context(), req)
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
		"message": "Successfully reset password user",
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

func (t *UserHandler) GetProfie(c *fiber.Ctx) error {
	id := middleware.UserID(c)
	res, er := t.UserUC.GetUserById(c.Context(), uint(id))
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

func (t *UserHandler) GetDashboardKemahasiswaan(c *fiber.Ctx) error {
	res, er := t.UserUC.PengajuanEventOrmawa(c.Context())
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
