package v1

import (
	"ctfm_backend/global"
	"ctfm_backend/global/response"
	"ctfm_backend/models"
	requ "ctfm_backend/models/request"
	resp "ctfm_backend/models/response"
	"ctfm_backend/services"
	"ctfm_backend/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddChallengeAPI(c *gin.Context) {
	var C requ.EditChallengeStruct
	_ = c.ShouldBindJSON(&C)
	global.CTFM_LOG.Debug(C)
	ChallengeVerify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Category":    {utils.NotEmpty()},
		"Points":      {utils.NotEmpty()},
		"Flag":        {utils.NotEmpty()},
	}
	ChallengeVerifyErr := utils.Verify(C, ChallengeVerify)

	if ChallengeVerifyErr != nil {
		response.FailWithMessage(ChallengeVerifyErr.Error(), c)
		return
	}

	challenge := &models.Challenge{Name: C.Name, Points: C.Points, Description: C.Description, Category: strings.ToLower(C.Category), Flag: C.Flag, IsHidden: C.IsHidden}
	err, id := services.AddChallenge(*challenge)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkDetailed(resp.ChallengedAddedResponse{ID: id}, "Challenge Added Successfully", c)
	}
}

func EditChallengeAPI(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var C requ.EditChallengeStruct
	_ = c.ShouldBindJSON(&C)
	ChallengeVerify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Category":    {utils.NotEmpty()},
		"Points":      {utils.NotEmpty()},
		"Flag":        {utils.NotEmpty()},
	}
	ChallengeVerifyErr := utils.Verify(C, ChallengeVerify)

	if ChallengeVerifyErr != nil {
		response.FailWithMessage(ChallengeVerifyErr.Error(), c)
		return
	}
	challenge := &models.Challenge{Name: C.Name, Points: C.Points, Description: C.Description, Category: strings.ToLower(C.Category), Flag: C.Flag, IsHidden: C.IsHidden}
	err := services.EditChallengeByID(*challenge, id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("Challenge Edited Successfully", c)
	}
}

// @Tags Challenges
// @Summary 获取题目列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param category query true "分类名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /challenges [get]
func GetChallengesListAPI(c *gin.Context) {
	category := c.Query("category")
	if category != "" {
		category = strings.ToLower(category)
		err, challenges := services.GetChallengesListByCategory(category)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
		} else {
			response.OkWithData(resp.ChallengesListResponse{Challenges: challenges}, c)
		}
		return
	}
	err, challenges := services.GetChallengesList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
	} else {
		response.OkWithData(resp.ChallengesListResponse{Challenges: challenges}, c)
	}
}

// @Tags Challenges
// @Summary 获取题目详细信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Param challengid param true "题目id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /challenges/:id [get]
func GetChallengeByIDAPI(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	with := c.Query("with")
	if with == "detail" {
		err, challenge := services.GetChallengeDetailByID(id)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
		} else {
			response.OkWithData(challenge, c)
		}
		return
	}
	err, challenge := services.GetChallengeContentByID(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
	} else {
		response.OkWithData(challenge, c)
	}
}

// @Tags Challenges
// @Summary 删除题目
// @Security ApiKeyAuth
// @Produce  application/json
// @Param challengid param true "题目id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /challenges/:id [DELETE]
func DeleteChallengeByIDAPI(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteChallengeByID(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
	} else {
		response.OkWithMessage("Delete Successfully", c)
	}
}

// @Tags Challenges
// @Summary 校验Flag
// @Security ApiKeyAuth
// @Produce  application/json
// @Param challengid param true "题目id"
// @Param data body {"flag":"xxx"} true "待校验flag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /challenges/:id/checkflag [GET]
func ValidChallengeFlagByIDAPI(c *gin.Context) {
	var F requ.CheckFlagStruct
	_ = c.ShouldBindJSON(&F)
	id, _ := strconv.Atoi(c.Param("id"))
	err, flag := services.GetChallengeFlagByID(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
	} else {
		if flag == F.Flag {
			response.OkWithMessage("Right Flag!", c)
		} else {
			response.FailWithMessage("Wrong Flag!", c)
		}
	}
}
