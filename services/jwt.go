package services

import (
	"ctfm_backend/global"
	"ctfm_backend/models"
)

// @title    JsonInBlacklist
// @description   create jwt blacklist
// @param     jwtList         model.JwtBlacklist
// @auth
// @return    err             error

func JsonInBlacklist(jwtList models.JwtBlacklist) (err error) {
	err = global.CTFM_DB.Create(&jwtList).Error
	return
}

// @title    IsBlacklist
// @description   check if the Jwt is in the blacklist or not, 判断JWT是否在黑名单内部
// @auth
// @param     jwt             string
// @param     jwtList         model.JwtBlacklist
// @return    err             error

func IsBlacklist(jwt string, jwtList models.JwtBlacklist) bool {
	isNotFound := global.CTFM_DB.Where("jwt = ?", jwt).First(&jwtList).RecordNotFound()
	return !isNotFound
}
