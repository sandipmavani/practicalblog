package controllers

import (
	"practicalblog/database"
	"practicalblog/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateBlogResponse struct {
	NickName string    `json:"nickname"`
	Id       uint      `json:"id"`
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
}

func CreateBlog(c *fiber.Ctx) error {

	var data map[string]string

	if error := c.BodyParser(&data); error != nil {
		return error
	}

	blogInfo := models.Blog{
		Title:    data["title"],
		Content:  data["content"],
		NickName: data["nickname"],
	}

	database.DB.Table("blog").Create(&blogInfo)

	var tmp CreateBlogResponse
	tmp.NickName = blogInfo.NickName
	tmp.Id = blogInfo.Id
	tmp.Title = blogInfo.Title
	return c.JSON(fiber.Map{
		"message": "success",
		"user":    tmp,
	})
}
