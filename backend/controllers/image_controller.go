package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = file.Filename
		if err := c.SaveFile(file, "./assets/uploads/"+filename); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(Response{
				Message: http.StatusText(http.StatusInternalServerError),
				Success: false,
				Status:  http.StatusInternalServerError,
			})
		}
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"url": "http://localhost:3000/api/v1/assets/uploads/" + filename,
		},
		Success: true,
		Status:  200,
	})
}
