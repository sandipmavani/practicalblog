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

	var articleId, err = c.ParamsInt("articleId")
	if err != nil {
		fmt.Println(err)
	}

	var result models.Article
	database.DB.Model(models.Article{Id: uint(articleId)}).First(&result)

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

func GetArticleComment(c *fiber.Ctx) error {

	var articleId, err = c.ParamsInt("articleId")
	if err != nil {
		fmt.Println(err)
	}

	var result []models.Comment
	database.DB.Model(models.Comment{BlogId: uint(articleId)}).Find(&result)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    result,
	})
}

func GetAllArticle(c *fiber.Ctx) error {

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var result []models.Article
	database.DB.Model(models.Article{}).Find(&result).Offset(offset).Limit(pageSize)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    result,
	})
}
