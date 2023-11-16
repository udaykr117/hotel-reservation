package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vUdayKumarr/hotel-reservation/types"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return fmt.Errorf("not authorised")
	}
	if !user.IsAdmin {
		return fmt.Errorf("not authorised")
	}
	return c.Next()
}
