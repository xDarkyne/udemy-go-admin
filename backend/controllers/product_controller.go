package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/db"
	"github.com/xdarkyne/udemy/db/models"
)

func AllProducts(c *fiber.Ctx) error {
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

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	if err := db.ORM.Create(&product).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Success: false,
			Status:  http.StatusInternalServerError,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"product": product,
		},
		Success: true,
		Status:  200,
	})
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product
	if err := db.ORM.First(&product, "product_id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"product": product,
		},
		Success: true,
		Status:  200,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	product := models.Product{
		ProductID: id,
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	if err := db.ORM.Model(&product).Updates(product).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"product": product,
		},
		Success: true,
		Status:  200,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := db.ORM.Delete(&models.Product{}, "product_id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Message: fmt.Sprintf("Product %s successfully deleted", id),
		Success: true,
		Status:  200,
	})
}
