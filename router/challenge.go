package router

import (
	v1 "ctfm_backend/api/v1"

	"github.com/gin-gonic/gin"
)

func InitChallengeRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("challenges")
	{
		UserRouter.POST("", v1.AddChallengeAPI)
		UserRouter.PUT(":id", v1.EditChallengeAPI)
		UserRouter.GET("", v1.GetChallengesListAPI)
		UserRouter.GET(":id", v1.GetChallengeByIDAPI)
		UserRouter.DELETE(":id", v1.DeleteChallengeByIDAPI)
		UserRouter.GET(":id/checkflag", v1.ValidChallengeFlagByIDAPI)
	}
}
