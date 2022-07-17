package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/config"
	"github.com/xdarkyne/udemy/controllers"
	"github.com/xdarkyne/udemy/db"
	"github.com/xdarkyne/udemy/db/models"
	"github.com/xdarkyne/udemy/util"
)

func WithAuth(c *fiber.Ctx) error {
	cookie := c.Cookies(config.App.AuthCookieName)

	issuer, err := util.ParseJWT(cookie)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(controllers.Response{
			Message: http.StatusText(http.StatusUnauthorized),
			Success: false,
			Status:  http.StatusUnauthorized,
		})
	}

	var user models.User
	if err := db.ORM.Preload("Role").First(&user, "user_id = ?", issuer).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(controllers.Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Success: false,
			Status:  http.StatusInternalServerError,
		})
	}

	c.Locals("user", user)
	return c.Next()
}
