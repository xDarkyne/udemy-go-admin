package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/db"
	"github.com/xdarkyne/udemy/db/models"
)

func AllUsers(c *fiber.Ctx) error {
	data, err := db.Paginate(c, db.ORM, &models.User{})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"users": data["data"],
			"meta":  data["meta"],
		},
		Success: true,
		Status:  200,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	user.HashPassword(user.Password)

	if err := db.ORM.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Success: false,
			Status:  http.StatusInternalServerError,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"user": user,
		},
		Success: true,
		Status:  200,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := db.ORM.Preload("Role").First(&user, "user_id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"user": user,
		},
		Success: true,
		Status:  200,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := models.User{
		UserID: id,
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	if err := db.ORM.Model(&user).Updates(user).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"user": user,
		},
		Success: true,
		Status:  200,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := db.ORM.Delete(&models.User{}, "user_id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Message: fmt.Sprintf("User %s successfully deleted", id),
		Success: true,
		Status:  200,
	})
}

func Export(c *fiber.Ctx) error {
	filePath := "./assets/csv/users.csv"
	if err := createFile(filePath); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Success: false,
			Status:  http.StatusInternalServerError,
		})
	}

	return c.Download(filePath)
}

func createFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var users []models.User
	if err := db.ORM.Preload("Role").Find(&users).Error; err != nil {
		return err
	}

	writer.Write([]string{
		"UserID", "Username", "Email",
	})

	for _, user := range users {
		data := []string{
			user.UserID,
			user.Username,
			user.Email,
		}

		if err := writer.Write(data); err != nil {
			return err
		}
	}

	return nil
}
