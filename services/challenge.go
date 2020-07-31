package services

import (
	"ctfm_backend/global"
	"ctfm_backend/models"
)

// @title    AddChallenge
// @description   AddChallenge, 新建题目
// @param     c               model.Challenge
// @return    err             error
func AddChallenge(c models.Challenge) (err error) {
	err = global.CTFM_DB.Create(&c).Error
	return err
}

func GetAllChallenges() (err error, challenges []models.Challenge) {
	err = global.CTFM_DB.Find(&challenges).Error
	return
}

/*func GetChallengesByCategory() (err error, challenges []models.Challenge) {

}*/
