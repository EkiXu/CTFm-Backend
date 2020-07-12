package router

import (
	v1 "ctfm_backend/api/v1"
	"ctfm_backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitJWTRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use(middleware.JWTAuth())
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
