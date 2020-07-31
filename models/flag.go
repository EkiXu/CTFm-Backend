package models

import "github.com/jinzhu/gorm"

type Flag struct {
	gorm.Model
	ChallengeID string `json:"challenge_id" gorm:"comment:'题目id'"`
	Content     string `json:"content" gorm:"comment:'内容'"`
}
