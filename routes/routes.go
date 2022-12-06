package routes

import (
	"practicalblog/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	//user operation
	app.Post("/api/blog/create", controllers.CreateBlog)
}
