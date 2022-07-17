package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/db"
	"github.com/xdarkyne/udemy/db/models"
)

func AllRoles(c *fiber.Ctx) error {
	data, err := db.Paginate(c, db.ORM, &models.Role{})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"roles": data["data"],
			"meta":  data["meta"],
		},
		Success: true,
		Status:  200,
	})
}

type RoleCreateDTO struct {
	Name        string `json:"name"`
	Permissions []uint `json:"permissions"`
}

func CreateRole(c *fiber.Ctx) error {
	var roleDTO RoleCreateDTO

	if err := c.BodyParser(&roleDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	permissions := make([]models.Permission, len(roleDTO.Permissions))

	for i, v := range roleDTO.Permissions {
		permissions[i] = models.Permission{
			Id: v,
		}
	}

	role := models.Role{
		Name:        roleDTO.Name,
		Permissions: permissions,
	}

	if err := db.ORM.Create(&role).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Success: false,
			Status:  http.StatusInternalServerError,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"role": role,
		},
		Success: true,
		Status:  200,
	})
}

func GetRole(c *fiber.Ctx) error {
	id := c.Params("id")

	var role models.Role
	if err := db.ORM.Preload("Permissions").First(&role, "role_id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"role": role,
		},
		Success: true,
		Status:  200,
	})
}

func UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")

	var roleDTO RoleCreateDTO

	if err := c.BodyParser(&roleDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	permissions := make([]models.Permission, len(roleDTO.Permissions))

	for i, v := range roleDTO.Permissions {
		permissions[i] = models.Permission{
			Id: v,
		}
	}

	role := models.Role{
		RoleID: id,
		Name:   roleDTO.Name,
	}

	if err := db.ORM.Model(&role).Association("Permissions").Clear(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Success: false,
			Status:  http.StatusInternalServerError,
		})
	}

	role.Permissions = permissions

	if err := db.ORM.Model(&role).Updates(role).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Data: fiber.Map{
			"role": role,
		},
		Success: true,
		Status:  200,
	})
}

func DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := db.ORM.Delete(&models.Role{}, "role_id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: http.StatusText(http.StatusNotFound),
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(Response{
		Message: fmt.Sprintf("Role %s successfully deleted", id),
		Success: true,
		Status:  200,
	})
}
