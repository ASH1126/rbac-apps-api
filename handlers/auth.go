package handlers

import (
	"fmt"
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
		response.Success = false
		response.Message = err.Error()
		return c.JSON(response)
	}

	fmt.Println(creds.UserName)
	fmt.Println(creds.Password)

	// email := c.FormValue("email")
	// pin := c.FormValue("pin")
	// data := handler.Login(email, pin)
	// fmt.Println(c.Send(c.Body()))

	return c.JSON(creds)
}
