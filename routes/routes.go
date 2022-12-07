package routes

import (
	"practicalblog/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	//user operation
	app.Post("/api/post/article", controllers.PostArticle)
	app.Post("/api/comment/article", controllers.CommentOnArticle)
	app.Post("/api/comment/comment", controllers.CommentOnComment)
	app.Get("/api/article/:articleId", controllers.GetArticleContent)

}
