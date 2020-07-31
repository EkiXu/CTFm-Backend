package services

import (
	"ctfm_backend/global"
	"ctfm_backend/models"
	"ctfm_backend/utils"
	"errors"

	uuid "github.com/satori/go.uuid"
)

// @title    Register
// @description   register, 用户注册
// @param     u               model.User
// @return    err             error
// @return    userInter       *User

func Register(u models.User) (err error, userInter models.User) {
	var user models.User
	// 判断用户名是否注册
	notRegister := global.CTFM_DB.Where("username = ?", u.Username).First(&user).RecordNotFound()
	// notRegister为false表明读取到了 不能注册
	if !notRegister {
		return errors.New("用户名已注册"), userInter
	} else {
		u.Password = utils.MSHA256([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.CTFM_DB.Create(&u).Error
	}
	return err, u
}

// @title    Login
// @description   login, 用户登录
// @param     u               *model.User
// @return    err             error
// @return    userInter       *User

func Login(u *models.User) (err error, userInter *models.User) {
	var user models.User
	u.Password = utils.MSHA256([]byte(u.Password))
	err = global.CTFM_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}
