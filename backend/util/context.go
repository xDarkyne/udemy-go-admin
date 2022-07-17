package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/db/models"
)

func GetUser(c *fiber.Ctx) models.User {
	user := c.Locals("user").(models.User)
	return user
}
