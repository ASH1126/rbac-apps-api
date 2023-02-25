package utils

import "github.com/gofiber/fiber/v2"

type Respon struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ErrorJSON(c *fiber.Ctx, messages string) error {
	var respon Respon
	respon.Success = false
	respon.Message = messages
	return c.JSON(respon)
}
