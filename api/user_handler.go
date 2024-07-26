package api

import "github.com/gofiber/fiber/v2"

func HandleListGetUser(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{
		"message": "Add the user",
	})
}
