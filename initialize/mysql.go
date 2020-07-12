package initialize

import (
	"ctfm_backend/global"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化数据库并产生数据库全局变量
func Mysql() {
	admin := global.CTFM_CONFIG.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		global.CTFM_LOG.Error("MySQL启动异常", err)
		os.Exit(0)
	} else {
		global.CTFM_DB = db
		global.CTFM_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.CTFM_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.CTFM_DB.LogMode(admin.LogMode)
	}
}
