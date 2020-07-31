package router

import (
	v1 "ctfm_backend/api/v1"

	"github.com/gin-gonic/gin"
)

func InitChallengeRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("challenges")
	{
		UserRouter.POST("", v1.AddChallengeAPI)
		UserRouter.GET("", v1.GetChallengesAPI)
	}
}
