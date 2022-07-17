package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/config"
	"github.com/xdarkyne/udemy/db"
	"github.com/xdarkyne/udemy/db/models"
	"github.com/xdarkyne/udemy/util"
)

type Response struct {
	Data    fiber.Map `json:"data,omitempty"`
	Message string    `json:"message,omitempty"`
	Status  int       `json:"status"`
	Success bool      `json:"success"`
}

type RegisterBody struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type UpdatePasswordBody struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.StandardClaims
}

func Register(c *fiber.Ctx) error {
	var data RegisterBody

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	if data.Password != data.PasswordConfirm {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: "Passwords do not match",
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	user := models.User{
		Username: data.Username,
		Email:    data.Email,
	}
	user.HashPassword(data.Password)

	if err := db.ORM.Create(&user).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: "Username is already taken",
			Success: false,
			Status:  http.StatusBadRequest,
		})

	}

	db.ORM.Find(&user, "username = ?", data.Username)

	return c.JSON(Response{
		Data: fiber.Map{
			"user": user,
		},
		Success: true,
		Status:  200,
	})
}

func Login(c *fiber.Ctx) error {
	var data LoginBody

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	var user models.User
	if err := db.ORM.First(&user, "username = ?", data.Username).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(Response{
			Message: "Could not find user",
			Success: false,
			Status:  http.StatusNotFound,
		})
	}

	if !user.IsValidPassword(data.Password) {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: "Incorrect credentials",
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	token, err := util.GenerateJWT(user.UserID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Success: false,
			Status:  http.StatusInternalServerError,
		})
	}

	cookie := fiber.Cookie{
		Name:     config.App.AuthCookieName,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
	}

	c.Cookie(&cookie)
	return c.JSON(Response{
		Data: fiber.Map{
			"token": token,
		},
		Success: true,
		Status:  200,
	})
}

func Status(c *fiber.Ctx) error {
	user := util.GetUser(c)

	return c.JSON(Response{
		Data: fiber.Map{
			"user": user,
		},
		Success: true,
		Status:  200,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     config.App.AuthCookieName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
	}

	c.Cookie(&cookie)
	return c.JSON(Response{
		Message: "Successfully logged out",
		Success: true,
		Status:  200,
	})
}

func UpdatePassword(c *fiber.Ctx) error {
	var data UpdatePasswordBody

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: http.StatusText(http.StatusBadRequest),
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	if data.Password != data.PasswordConfirm {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Message: "Passwords do not match",
			Success: false,
			Status:  http.StatusBadRequest,
		})
	}

	user := util.GetUser(c)
	user.HashPassword(data.Password)

	if err := db.ORM.Model(&user).Where("user_id = ?", user.UserID).Updates(user).Error; err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Message: http.StatusText(http.StatusInternalServerError),
			Success: false,
			Status:  http.StatusInternalServerError,
		})
	}

	return c.JSON(Response{
		Message: "Successfully updated password",
		Success: true,
		Status:  200,
	})
}
