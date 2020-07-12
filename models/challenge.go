package models

import (
	"github.com/jinzhu/gorm"
)

type Challenge struct {
	gorm.Model
	Description string `json:"description" gorm:"comments:'题目描述'"`
	Category    string `json:"category" gorm:"comments:'题目分类'"`
	Flag        string `json:"flag" gorm:"comments:'题目答案'"`
	IsHidden    bool   `json:"IsAdmin" gorm:"comment:'是否隐藏';default:1"`
}
