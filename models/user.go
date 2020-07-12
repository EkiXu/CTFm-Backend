package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:'用户UUID'"`
	Username string    `json:"username" gorm:"comment:'用户登录名'"`
	Password string    `json:"-"  gorm:"comment:'用户登录密码'"`
	Email    string    `json:"email"  gorm:"comment:'用户邮箱'"`
	Nickname string    `json:"nickname" gorm:"comment:'用户昵称'" `
	IsAdmin  bool      `json:"IsAdmin" gorm:"comment:'是否管理员';default:0"`
}
