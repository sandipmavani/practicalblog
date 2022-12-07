package models

import "time"

type Article struct {
	Id       uint      `json:"id"`
	Title    string    `json:"title" gorm:"column:title"`
	Content  string    `json:"content" gorm:"column:content"`
	NickName string    `json:"nickname" gorm:"column:nickname"`
	Date     time.Time `json:"date" gorm:"column:date"`
}
