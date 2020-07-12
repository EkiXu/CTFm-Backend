package router

import (
	v1 "ctfm_backend/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", v1.LoginAPI)
		UserRouter.POST("register", v1.RegisterAPI) // 用户注册
	}
}
