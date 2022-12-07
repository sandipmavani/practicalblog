package controllers

import (
	"fmt"
	"practicalblog/database"
	"practicalblog/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateArticleResponse struct {
	NickName string    `json:"nickname"`
	Id       uint      `json:"id"`
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
}

type CreateCommentRequest struct {
	BlogID   int    `json:"blogId"`
	Content  string `json:"content"`
	Nickname string `json:"nickname"`
	ParentId int    `json:"parentId"`
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
	database.DB.Table("articles").Create(&articleInfo)

	var tmp CreateArticleResponse
	tmp.NickName = articleInfo.NickName
	tmp.Id = articleInfo.Id
	tmp.Title = articleInfo.Title
	tmp.Date = articleInfo.Date
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    tmp,
	})
}

func GetArticleContent(c *fiber.Ctx) error {

	var articleId = c.Params("articleId")
	u64, err := strconv.ParseUint(articleId, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	id := uint(u64)

	var result models.Article
	database.DB.Model(models.Article{Id: id}).First(&result)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    result,
	})
}

func CommentOnArticle(c *fiber.Ctx) error {

	var data CreateCommentRequest

	if error := c.BodyParser(&data); error != nil {
		return error
	}

	commentObj := models.Comment{
		BlogId:   uint(data.BlogID),
		Content:  data.Content,
		NickName: data.Nickname,
	}
	commentObj.Date = time.Now()
	database.DB.Table("comments").Create(&commentObj)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    commentObj,
	})
}

func CommentOnComment(c *fiber.Ctx) error {

	var data CreateCommentRequest

	if error := c.BodyParser(&data); error != nil {
		return error
	}

	commentObj := models.Comment{
		BlogId:   uint(data.BlogID),
		Content:  data.Content,
		NickName: data.Nickname,
		ParentId: uint(data.ParentId),
	}
	commentObj.Date = time.Now()
	database.DB.Table("comments").Create(&commentObj)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    commentObj,
	})
}
