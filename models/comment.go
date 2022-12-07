package models

import "time"

type Comment struct {
	Id       uint      `json:"id"`
	ParentId uint      `json:"parentId" gorm:"column:parentId"`
	ArticleId   uint      `json:"articleId" gorm:"column:articleId"`
	Content  string    `json:"content" gorm:"column:content"`
	NickName string    `json:"nickname" gorm:"column:nickname"`
	Date     time.Time `json:"date" gorm:"column:date"`
}
