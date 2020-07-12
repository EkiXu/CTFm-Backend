package initialize

import (
	"ctfm_backend/global"
	"ctfm_backend/models"
)

// 注册数据库表专用
func DBTables() {
	db := global.CTFM_DB
	db.AutoMigrate(models.User{},
		models.JwtBlacklist{})
	global.CTFM_LOG.Debug("register table success")
}
