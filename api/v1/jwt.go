package v1

import (
	"ctfm_backend/global/response"
	"ctfm_backend/models"
	"ctfm_backend/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"jwt作废成功"}"
// @Router /jwt/jsonInBlacklist [post]
func JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	modelJwt := models.JwtBlacklist{
		Jwt: token,
	}
	err := services.JsonInBlacklist(modelJwt)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("jwt作废失败，%v", err), c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
