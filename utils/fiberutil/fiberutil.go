package fiberutil

import "github.com/gofiber/fiber/v2"

func ReturnStatusUnprocessableEntity(c *fiber.Ctx, messages any, errorData any) error {
	statusCode := fiber.StatusUnprocessableEntity
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"status":  statusCode,
		"message": messages,
		// "error":   errorData,
		"data": map[string]map[string]any{
			"errors": {
				"general": messages,
			},
			// []map[string]any{}
			// 	"errors": any,
			// },
		},
	})
}
func ReturnStatusBadRequest(c *fiber.Ctx, messages any, errorData any) error {
	statusCode := fiber.StatusBadRequest
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"status":  statusCode,
		"message": messages,
		// "error":   errorData,
		"data": map[string]map[string]any{
			"errors": {
				"general": messages,
			},
			// []map[string]any{}
			// 	"errors": any,
			// },
		},
	})
}

func ReturnErrorCustomStatusType(c *fiber.Ctx, messages any, errorData any, statusCode int, errorType string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"status":  statusCode,
		"message": messages,
		"data": map[string]map[string]any{
			"errors": {
				errorType: messages,
			},
		},
	})
}

func ReturnStatusNotFound(c *fiber.Ctx, messages []string, errorData any) error {
	statusCode := fiber.StatusNotFound
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"status":  statusCode,
		"message": messages,
		"error":   errorData,
	})
}

func ReturnStatusOK(c *fiber.Ctx, messages []string, data any) error {
	statusCode := fiber.StatusOK
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"status":  statusCode,
		"message": messages,
		"data":    data,
	})
}

func ReturnStatusUnauthorized(c *fiber.Ctx) error {
	statusCode := fiber.StatusPaymentRequired
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"status":  statusCode,
		"message": []string{"Unauthorized"},
		"data":    []string{},
	})
}

func ReturnStatusUnauthorizedError(c *fiber.Ctx, err error, errorType string) error {
	statusCode := fiber.StatusPaymentRequired
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"status":  statusCode,
		"message": []string{"Unauthorized"},
		"data": map[string]map[string]any{
			"errors": {
				errorType: err.Error(),
			},
		},
	})
}

func ReturnStatusOKLegacy(c *fiber.Ctx, messages []string, data any) error {
	statusCode := fiber.StatusOK
	return c.Status(statusCode).JSON(fiber.Map{
		"success":  true,
		"status":   statusCode,
		"messages": messages,
		"meta":     []string{},
		"data":     data,
	})
}
func ReturnStatusCreatedLegacy(c *fiber.Ctx, messages []string, data any) error {
	statusCode := fiber.StatusCreated
	return c.Status(statusCode).JSON(fiber.Map{
		"success":  true,
		"status":   statusCode,
		"messages": messages,
		"meta":     []string{},
		"data":     data,
	})
}

func ReturnStatusUnprocessableEntityLegacy(c *fiber.Ctx, messages any, errorData any) error {
	statusCode := fiber.StatusUnprocessableEntity
	return c.Status(statusCode).JSON(fiber.Map{
		"success":  false,
		"status":   statusCode,
		"messages": messages,
		"meta":     []string{},
		// "error":   errorData,
		"data": map[string]map[string]any{
			"errors": {
				"general": messages,
			},
			// []map[string]any{}
			// 	"errors": any,
			// },
		},
	})
}
