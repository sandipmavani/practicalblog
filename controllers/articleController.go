package controllers

import (
	"practicalblog/database"
	"practicalblog/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateArticleResponse struct {
	NickName string    `json:"nickname"`
	Id       uint      `json:"id"`
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
}

func PostArticle(c *fiber.Ctx) error {

	var data map[string]string

	if error := c.BodyParser(&data); error != nil {
		return error
	}

	articleInfo := models.Article{
		Title:    data["title"],
		Content:  data["content"],
		NickName: data["nickname"],
	}
	articleInfo.Date = time.Now()
	database.DB.Table("blogs").Create(&articleInfo)

	var tmp CreateArticleResponse
	tmp.NickName = articleInfo.NickName
	tmp.Id = articleInfo.Id
	tmp.Title = articleInfo.Title
	tmp.Date = articleInfo.Date
	return c.JSON(fiber.Map{
		"message": "success",
		"user":    tmp,
	})
}