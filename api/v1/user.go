package v1

import (
	"ctfm_backend/global"
	"ctfm_backend/global/response"
	"ctfm_backend/middleware"
	"ctfm_backend/models"
	requ "ctfm_backend/models/request"
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
	var R requ.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	UserVerify := utils.Rules{
		"Email":    {utils.NotEmpty()},
		"Username": {utils.NotEmpty()},
		"Nickname": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)
	if !utils.IsValidEmail(R.Email) && !utils.IsValidName(R.Nickname) && !utils.IsValidName(R.Username) {
		response.FailWithMessage("Format Invalid", c)
		return
	}
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}

	user := &models.User{Username: R.Username, Nickname: R.Nickname, Email: R.Email, Password: R.Password}
	err, userReturn := services.Register(*user)
	if err != nil {
		response.FailWithCodeAndMessage(response.ERROR, fmt.Sprintf("%v", err), c)
	} else {
		response.OkDetailed(userReturn, "Register Successfully", c)
	}
}

// @Tags User
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.LoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /user/login [post]
func LoginAPI(c *gin.Context) {
	var L requ.LoginStruct
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
	if !utils.IsValidName(L.Username) {
		response.FailWithMessage("Invalid Username", c)
		return
	}
	U := &models.User{Username: L.Username, Password: L.Password}
	if err, user := services.Login(U); err != nil {
		response.FailWithMessage("Username or Password is Wrong ", c)
	} else {
		tokenNext(c, *user)
	}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user models.User) {
	j := &middleware.JWT{
		SigningKey: []byte(global.CTFM_CONFIG.JWT.SigningKey), // 唯一签名
	}
	claims := requ.CustomClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		Nickname: user.Nickname,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "CTFm",                         // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("Failed to get token", c)
		return
	}
	response.OkWithData(resp.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, c)
	return
}

/*func AddUserAPI(c *gin.Context) {
	var C requ.EditUserStruct
	_ = c.ShouldBindJSON(&C)
	global.CTFM_LOG.Debug(C)
	UserVerify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Category":    {utils.NotEmpty()},
		"Points":      {utils.NotEmpty()},
		"Flag":        {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(C, UserVerify)

	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}

	user := &models.User{Name: C.Name, Points: C.Points, Description: C.Description, Category: strings.ToLower(C.Category), Flag: C.Flag, IsHidden: C.IsHidden}
	err, id := services.AddUser(*user)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkDetailed(resp.UserdAddedResponse{ID: id}, "User Added Successfully", c)
	}
}

func EditUserAPI(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var C requ.EditUserStruct
	_ = c.ShouldBindJSON(&C)
	UserVerify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Category":    {utils.NotEmpty()},
		"Points":      {utils.NotEmpty()},
		"Flag":        {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(C, UserVerify)

	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	user := &models.User{Name: C.Name, Points: C.Points, Description: C.Description, Category: strings.ToLower(C.Category), Flag: C.Flag, IsHidden: C.IsHidden}
	err := services.EditUserByID(*user, id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("User Edited Successfully", c)
	}
}

// @Tags Users
// @Summary 获取用户列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param category query true "分类名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /users [get]
func GetUsersListAPI(c *gin.Context) {
	category := c.Query("category")
	if category != "" {
		category = strings.ToLower(category)
		err, users := services.GetUsersListByCategory(category)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
		} else {
			response.OkWithData(resp.UsersListResponse{Users: users}, c)
		}
		return
	}
	err, users := services.GetUsersList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
	} else {
		response.OkWithData(resp.UsersListResponse{Users: users}, c)
	}
}

// @Tags Users
// @Summary 获取用户详细信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Param challengid param true "用户id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /users/:id [get]
func GetUserByIDAPI(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	with := c.Query("with")
	if with == "detail" {
		err, user := services.GetUserDetailByID(id)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
		} else {
			response.OkWithData(user, c)
		}
		return
	}
	err, user := services.GetUserContentByID(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
	} else {
		response.OkWithData(user, c)
	}
}

// @Tags Users
// @Summary 删除用户
// @Security ApiKeyAuth
// @Produce  application/json
// @Param challengid param true "用户id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /users/:id [DELETE]
func DeleteUserByIDAPI(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteUserByID(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Fail，%v", err), c)
	} else {
		response.OkWithMessage("Delete Successfully", c)
	}
}
*/
