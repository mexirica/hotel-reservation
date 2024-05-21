package api

import (
	"github.com/mexirica/hotel-reservation/types"
	"github.com/gofiber/fiber/v3"
)

func AdminAuth(c fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return ErrUnAuthorized()
	}
	if !user.IsAdmin {
		return ErrUnAuthorized()
	}
	return c.Next()
}
