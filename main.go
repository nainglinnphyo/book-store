package main

import (
	"bookapi.com/config"
	"bookapi.com/database"

	_ "bookapi.com/docs"
	"bookapi.com/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	database.Connect()
	p := config.Config("APP_PORT")
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	api := app.Group("api")
	routes.MainRouteGroup(api)
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Listen(p)
}
