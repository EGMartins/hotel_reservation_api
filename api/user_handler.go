package api

import (
	"github.com/egmartins/hotel_rservation_api/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
    FistName: "Eduardo",
    LastName: "Martins",
  }
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
