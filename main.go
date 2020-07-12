package main

import (
	"ctfm_backend/core"
	"ctfm_backend/global"
	"ctfm_backend/initialize"
)

// @title CTFm API
// @version 0.0.1
// @description  This is the backend for CTFm.
// @BasePath /api/v1/
func main() {
	initialize.Mysql()
	initialize.DBTables()
	// 程序结束前关闭数据库链接
	defer global.CTFM_DB.Close()
	core.RunServer()
}
