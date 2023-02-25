package controllers

import (
	"rbac-api/handlers"
	"rbac-api/utils"

	"github.com/gofiber/fiber/v2"
)

var response utils.Respon

func TestCall(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	response.Success = true
	response.Message = "Welcome to SERVICE API"
	response.Data = data
	return c.JSON(response)
}

func Login(c *fiber.Ctx) error {
	type credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
	}
	var creds credentials
	if err := c.BodyParser(&creds); err != nil {
		return utils.ErrorJSON(c, err.Error())
	}
	data := handlers.Login(creds.UserName, creds.Password)
	return c.JSON(data)
}
