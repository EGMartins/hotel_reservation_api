package api

import (
	"context"

	"github.com/egmartins/hotel_rservation_api/db"
	"github.com/egmartins/hotel_rservation_api/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandle struct {
	userStore db.UserStore
}

func NewUserHandle(userStore db.UserStore) *UserHandle {
	return &UserHandle{
		userStore: userStore,
	}
}

func (h *UserHandle) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandle) HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Eduardo",
		LastName:  "Martins",
	}
	return c.JSON(user)
}
