package models

import (
	"github.com/jinzhu/gorm"
)

type Challenge struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description" gorm:"comments:'题目描述'"`
	Category    string `json:"category" gorm:"comments:'题目分类'"`
	Points      int    `json:"points" gorm:"comments:'题目分值'"`
	IsHidden    bool   `json:"is_hidden" gorm:"comment:'是否隐藏';default:true"`
	Solved      int    `json:"solved" gorm:"comments:'解出人数';default:0"`
	Attempts    int    `json:"attempts" gorm:"comments:'尝试次数';default:0"`
	Flag string `json:"flag" gorm:"comments:'Flag'"`
}
