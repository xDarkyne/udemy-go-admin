package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/controllers"
	"github.com/xdarkyne/udemy/db"
	"github.com/xdarkyne/udemy/util"
)

func WithPermission(c *fiber.Ctx) error {
	user := util.GetUser(c)

	page := c.Locals("page").(string)

	db.ORM.Preload("Permissions").Find(&user.Role)

	hasRole := false

	if c.Method() == "GET" {
		for _, permission := range user.Role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				hasRole = true
			}
		}
	} else {
		for _, permission := range user.Role.Permissions {
			if permission.Name == "edit_"+page {
				hasRole = true
			}
		}
	}

	if !hasRole {
		return c.Status(http.StatusUnauthorized).JSON(controllers.Response{
			Message: http.StatusText(http.StatusUnauthorized),
			Success: false,
			Status:  http.StatusUnauthorized,
		})
	}

	return c.Next()
}
