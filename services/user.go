package services

import (
	"ctfm_backend/global"
	"ctfm_backend/models"
	resp "ctfm_backend/models/response"
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

// @title    AddUser
// @description   AddUser, 新建用户
// @param     c               model.User
// @return    err             error
func AddUser(c models.User) (err error, id uint) {
	err = global.CTFM_DB.Create(&c).Error
	if c.IsAdmin == false {
		global.CTFM_DB.Model(&c).Select("is_admin").Updates(map[string]interface{}{"is_hidden": false})
	}
	id = c.ID
	return
}

func EditUserByID(c models.User, id int) (err error) {
	var o models.User
	global.CTFM_DB.First(&o, id)
	err = global.CTFM_DB.Model(&o).Update(&c).Error
	if c.IsAdmin == false {
		global.CTFM_DB.Model(&c).Select("is_admin").Updates(map[string]interface{}{"is_hidden": false})
	}
	return err
}

func DeleteUserByID(id int) (err error) {
	var o models.User
	global.CTFM_DB.First(&o, id)
	err = global.CTFM_DB.Unscoped().Delete(&o).Error
	return err
}

func GetUsersList() (err error, users []resp.UserResponse) {
	err = global.CTFM_DB.Table("users").Select("id,username,nickname,score,solved,is_admin").Scan(&users).Error
	return
}
