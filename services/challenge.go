package services

import (
	"ctfm_backend/global"
	"ctfm_backend/models"
	resp "ctfm_backend/models/response"
)

// @title    AddChallenge
// @description   AddChallenge, 新建题目
// @param     c               model.Challenge
// @return    err             error
func AddChallenge(c models.Challenge) (err error, id uint) {
	err = global.CTFM_DB.Create(&c).Error
	if c.IsHidden == false {
		global.CTFM_DB.Model(&c).Select("is_hidden").Updates(map[string]interface{}{"is_hidden": false})
	}
	id = c.ID
	return
}

func EditChallengeByID(c models.Challenge, id int) (err error) {
	var o models.Challenge
	global.CTFM_DB.First(&o, id)
	err = global.CTFM_DB.Model(&o).Update(&c).Error
	if c.IsHidden == false {
		global.CTFM_DB.Model(&c).Select("is_hidden").Updates(map[string]interface{}{"is_hidden": false})
	}
	return err
}

func DeleteChallengeByID(id int) (err error) {
	var o models.Challenge
	global.CTFM_DB.First(&o, id)
	err = global.CTFM_DB.Unscoped().Delete(&o).Error
	return err
}

func GetChallengesList() (err error, challenges []resp.ChallengeResponse) {
	err = global.CTFM_DB.Table("challenges").Select("id,name,category,solved,attempts,points,is_hidden").Scan(&challenges).Error
	return
}

func GetChallengesListByCategory(category string) (err error, challenges []resp.ChallengeResponse) {
	err = global.CTFM_DB.Table("challenges").Select("id,name,category,solved,attempts,points").Not("is_hidden", 1).Where("category = ?", category).Scan(&challenges).Error
	return
}

func GetChallengeContentByID(id int) (err error, challenge resp.ChallengeContentResponse) {
	err = global.CTFM_DB.Table("challenges").Select("id,name,category,description,points").Where("id = ?", id).Scan(&challenge).Error
	return
}

func GetChallengeDetailByID(id int) (err error, challenge resp.ChallengeDetailResponse) {
	err = global.CTFM_DB.Table("challenges").Select("id,name,category,description,points,flag,is_hidden").Where("id = ?", id).Scan(&challenge).Error
	return
}
func GetChallengeFlagByID(id int) (err error, flag string) {
	err = global.CTFM_DB.Table("challenges").Select("flag").Where("id = ?", id).Scan(&flag).Error
	return
}
