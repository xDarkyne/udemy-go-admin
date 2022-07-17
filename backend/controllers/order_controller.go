package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/db"
	"github.com/xdarkyne/udemy/db/models"
)

func AllOrders(c *fiber.Ctx) error {
	data, err := db.Paginate(c, db.ORM, &models.Order{})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"orders": data["data"],
			"meta":   data["meta"],
		},
		Success: true,
		Status:  200,
	})
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	if err := db.ORM.Create(&order).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Success: false,
			Status:  http.StatusInternalServerError,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"order": order,
		},
		Success: true,
		Status:  200,
	})
}

func GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order
	if err := db.ORM.First(&order, "order_id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"order": order,
		},
		Success: true,
		Status:  200,
	})
}

func UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	order := models.Order{
		OrderID: id,
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	if err := db.ORM.Model(&order).Updates(order).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"order": order,
		},
		Success: true,
		Status:  200,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := db.ORM.Delete(&models.Order{}, "order_id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Message: fmt.Sprintf("Order %s successfully deleted", id),
		Success: true,
		Status:  200,
	})
}
