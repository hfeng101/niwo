package register

import (
	"github.com/kataras/iris/v12"
	"github.com/hfeng101/niwo/apis/user/v1"
)

//用户信息设置，含登陆、退出
func LoginRouteRegister(app *iris.Application) {
	loginParty := app.Party("/v1/login")

	//获取验证码
	loginParty.Post("/getVerificationCode", v1.GetVerificationCode)

	//验证码登陆
	loginParty.Post("/", v1.Login)

	//通过token登陆
	loginParty.Post("/verifyLogin", v1.VerifyLogin)

	//更新验证码
	loginParty.Post("/updateToken", v1.UpdateToken)

	//手机登陆
	//loginParty.Post("/phone",)

	//邮箱登陆
	//loginParty.Post("/email", )
}

func LogoutRouteRegister(app *iris.Application) {
	logoutParty := app.Party("/v1/logout")

	logoutParty.Post("/", v1.Logout)
}

func UserRouteRegister(app *iris.Application) {
	//userParty := app.Party("/user")

	//userParty.Get("/info",)
	//
	////修改密码
	//userParty.Put("/password",)
	//
	////添加、修改用户信息
	//userParty.Post("/info", )
	//userParty.Put("/info", )

}