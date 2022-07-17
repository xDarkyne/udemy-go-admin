package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/db"
	"github.com/xdarkyne/udemy/db/models"
)

func AllPermission(c *fiber.Ctx) error {
	data, err := db.Paginate(c, db.ORM, &models.Product{})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"products": data["data"],
			"meta":     data["meta"],
		},
		Success: true,
		Status:  200,
	})
}
