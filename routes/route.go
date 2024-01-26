package routes

import (
	"bookapi.com/controller"
	"github.com/gofiber/fiber/v2"
)

func MainRouteGroup(router fiber.Router) {
	router.Get("book", controller.GetAllBook)
	router.Post("book", controller.Create)
}
