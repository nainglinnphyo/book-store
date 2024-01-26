package routes

import (
	"bookapi.com/controller"
	"github.com/gofiber/fiber/v2"
)

func MainRouteGroup(router fiber.Router) {
	// @Param request body main.MyHandler.request true "query params"
	// @Success 200 {object} main.MyHandler.response
	// @Router /test [post]
	router.Get("book", controller.GetAllBook)
	router.Post("book", controller.Create)
}
