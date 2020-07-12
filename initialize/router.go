package initialize

import (
	_ "ctfm_backend/docs"
	"ctfm_backend/global"
	"ctfm_backend/router"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var r = gin.Default()

	r.LoadHTMLGlob("../dist/*.html")                                                            // 添加入口index.html
	r.Static("/static", "../dist/static")                                                       // 添加资源路径
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "PRODUCTION")) // 文档接口
	global.CTFM_LOG.Debug("register swagger handler")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// 方便统一添加路由组前缀 多服务器上线使用
	apiGroup := r.Group("/api/v1")

	router.InitUserRouter(apiGroup) // 注册用户路由
	router.InitJWTRouter(apiGroup)  // 注册JWT路由
	global.CTFM_LOG.Info("router register success")

	return r
}
