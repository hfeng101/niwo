package v1

import (
	UserController "github.com/hfeng101/niwo/controller/user/v1"
	"github.com/kataras/iris/v12"
)

// @title user operation API
// @version 1.0
// @contact.name user operation
// @contact.url http://www.niwofly.com
// @Description 用户登录，获取验证码，若是第一次登陆则自动注册
// @Accept  json
// @Produce  json
// @Param param body UserController.GetVerificationCodeReq true "param"
// @Success 20000 {object} UserController.ResponseContent	""
// @Failure 50001 {object} UserController.ResponseContent	""
// @Router /login/getVerificationCode [post]
// @BasePath /v1
func GetVerificationCode(ctx iris.Context) {
	UserController.GetVerificationCode(ctx)
}

// @title user operation API
// @version 1.0
// @contact.name user operation
// @contact.url http://www.niwofly.com
// @Description 获取到验证码后，验证登录
// @Accept  json
// @Produce  json
// @Param param body UserController.UserRegistrationOrLoginReq true "param"
// @Success 20000 {object} UserController.ResponseContent	""
// @Failure 50001 {object} UserController.ResponseContent	""
// @Router /login [post]
// @BasePath /v1
func Login(ctx iris.Context) {
	UserController.Login(ctx)
}

// @title user operation API
// @version 1.0
// @contact.name user operation
// @contact.url http://www.niwofly.com
// @Description 用户退出，清除所有登录状态
// @Accept  json
// @Produce  json
// @Param param body UserController.LogoutReq true "param"
// @Success 20000 {object} UserController.ResponseContent	""
// @Failure 50001 {object} UserController.ResponseContent	""
// @Router /logout [post]
// @BasePath /v1
func Logout(ctx iris.Context) {
	UserController.Logout(ctx)
}

// @title user operation API
// @version 1.0
// @contact.name user operation
// @contact.url http://www.niwofly.com
// @Description 验证用户以前登陆所生成的token来验证登陆
// @Accept  json
// @Produce  json
// @Param param body UserController.VerifyLoginReq true "param"
// @Success 20000 {object} UserController.ResponseContent	""
// @Failure 50001 {object} UserController.ResponseContent	""
// @Router /login/verifyLogin [post]
// @host niwofly.com
// @BasePath /v1
func VerifyLogin(ctx iris.Context) {
	UserController.VerifyLogin(ctx)
}

// @title user operation API
// @version 1.0
// @contact.name user operation
// @contact.url http://www.niwofly.com
// @Description 用户token过期，重复更新
// @Accept  json
// @Produce  json
// @Param param body UserController.UpdateTokenReq true "param"
// @Success 20000 {object} UserController.ResponseContent	""
// @Failure 50001 {object} UserController.ResponseContent	""
// @Router /login/updateToken [post]
// @host niwofly.com
// @BasePath /v1
func UpdateToken(ctx iris.Context) {
	UserController.UpdateToken(ctx)
}