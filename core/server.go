package core

import (
	"ctfm_backend/global"
	"ctfm_backend/initialize"
	"fmt"
	"net/http"
	"time"
)

func RunServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.CTFM_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.CTFM_LOG.Debug("server run success on ", address)

	global.CTFM_LOG.Error(s.ListenAndServe())
}
