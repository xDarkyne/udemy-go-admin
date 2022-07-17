package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdarkyne/udemy/controllers"
	"github.com/xdarkyne/udemy/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Use(middleware.WithAuth)
	auth.Get("/status", controllers.Status)
	auth.Put("/update", controllers.UpdatePassword)
	auth.Post("/logout", controllers.Logout)

	users := api.Group("/users", func(c *fiber.Ctx) error {
		c.Locals("page", "users")
		return c.Next()
	})
	users.Use(middleware.WithAuth)
	users.Use(middleware.WithPermission)
	users.Get("/", controllers.AllUsers)
	users.Post("/", controllers.CreateUser)
	users.Get("/:id", controllers.GetUser)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)
	users.Post("/export", controllers.Export)

	roles := api.Group("/roles", func(c *fiber.Ctx) error {
		c.Locals("page", "roles")
		return c.Next()
	})
	roles.Use(middleware.WithAuth)
	roles.Use(middleware.WithPermission)
	roles.Get("/", controllers.AllRoles)
	roles.Post("/", controllers.CreateRole)
	roles.Get("/:id", controllers.GetRole)
	roles.Put("/:id", controllers.UpdateRole)
	roles.Delete("/:id", controllers.DeleteRole)

	permissions := api.Group("/permissions", func(c *fiber.Ctx) error {
		c.Locals("page", "permissions")
		return c.Next()
	})
	permissions.Use(middleware.WithAuth)
	permissions.Use(middleware.WithPermission)
	permissions.Get("/", controllers.AllPermission)

	products := api.Group("/products", func(c *fiber.Ctx) error {
		c.Locals("page", "products")
		return c.Next()
	})
	products.Use(middleware.WithAuth)
	products.Use(middleware.WithPermission)
	products.Get("/", controllers.AllProducts)
	products.Post("/", controllers.CreateProduct)
	products.Get("/:id", controllers.GetProduct)
	products.Put("/:id", controllers.UpdateProduct)
	products.Delete("/:id", controllers.DeleteProduct)

	orders := api.Group("/orders", func(c *fiber.Ctx) error {
		c.Locals("page", "orders")
		return c.Next()
	})
	orders.Use(middleware.WithAuth)
	orders.Use(middleware.WithPermission)
	orders.Get("/", controllers.AllOrders)
	orders.Post("/", controllers.CreateOrder)
	orders.Get("/:id", controllers.GetOrder)
	orders.Put("/:id", controllers.UpdateOrder)
	orders.Delete("/:id", controllers.DeleteOrder)

	images := api.Group("/images", func(c *fiber.Ctx) error {
		c.Locals("page", "images")
		return c.Next()
	})
	images.Use(middleware.WithAuth)
	images.Post("/upload", controllers.Upload)

	api.Static("/assets", "./assets")
}
