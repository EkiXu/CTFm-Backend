package v1

import (
	"ctfm_backend/global"
	"ctfm_backend/global/response"
	"ctfm_backend/middleware"
	"ctfm_backend/models"
	"ctfm_backend/models/request"
	resp "ctfm_backend/models/response"
	"ctfm_backend/services"
	"ctfm_backend/utils"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Tags User
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.RegisterStruct true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Register Successfully"}"
// @Router /user/register [post]
func RegisterAPI(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	UserVerify := utils.Rules{
		"Email":    {utils.NotEmpty()},
		"Username": {utils.NotEmpty()},
		"Nickname": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)

	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}

	user := &models.User{Username: R.Username, Nickname: R.Nickname, Email: R.Email, Password: R.Password}
	err, userReturn := services.Register(*user)
	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.UserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		response.OkDetailed(resp.UserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags User
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.LoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /user/login [post]
func LoginAPI(c *gin.Context) {
	var L request.LoginStruct
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	U := &models.User{Username: L.Username, Password: L.Password}
	if err, user := services.Login(U); err != nil {
		response.FailWithMessage("username or password is wrong ", c)
	} else {
		tokenNext(c, *user)
	}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user models.User) {
	j := &middleware.JWT{
		SigningKey: []byte(global.CTFM_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := request.CustomClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		Nickname: user.Nickname,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "CTFm",                         // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithData(resp.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
	}, c)
	return
}
