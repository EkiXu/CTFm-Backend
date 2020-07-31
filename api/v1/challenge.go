package v1

import (
	"ctfm_backend/global/response"
	"ctfm_backend/models"
	requ "ctfm_backend/models/request"
	resp "ctfm_backend/models/response"
	"ctfm_backend/services"
	"ctfm_backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddChallengeAPI(c *gin.Context) {
	var C requ.AddChallengeStruct
	_ = c.ShouldBindJSON(&C)
	ChallengeVerify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Category":    {utils.NotEmpty()},
		"Points":      {utils.NotEmpty()},
	}
	ChallengeVerifyErr := utils.Verify(C, ChallengeVerify)

	if ChallengeVerifyErr != nil {
		response.FailWithMessage(ChallengeVerifyErr.Error(), c)
		return
	}

	challenge := &models.Challenge{Name: C.Name, Points: C.Points, Description: C.Description, Category: C.Category}
	err := services.AddChallenge(*challenge)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("题目添加成功", c)
	}
}

// @Tags authorityAndMenu
// @Summary 获取题目列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /challenges [get]
func GetChallengesAPI(c *gin.Context) {
	err, challenges := services.GetAllChallenges()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.OkWithData(resp.ChallengesResponse{Challenges: challenges}, c)
	}
}
