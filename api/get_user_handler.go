package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rishabh2030/hotel-reservation-system/entity"
)

func GetUserHandler(ctx *fiber.Ctx) error {
	user := entity.User{}
	user.ID = "1"
	user.FirstName = "Rishabh"
	user.LastName = "sharma"

	return ctx.JSON(user)
}
