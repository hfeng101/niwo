package v1

import (
	"github.com/kataras/iris/v12"
	UserController "github.com/hfeng101/niwo/controller/user/v1"
)

// @Description
// @Accept  json
// @Produce  json
// @Param param body UserController.UserRegistrationReq true "param"
// @Success 20000 {object} UserController.ResponseContent	""
// @Failure 50001 {object} UserController.ResponseContent	""
// @Router /login [post]
func Login(ctx iris.Context) {
	UserController.Login(ctx)
}

// @Description
// @Accept  json
// @Produce  json
// @Param param body UserController.UserRegistrationReq true "param"
// @Success 20000 {object} UserController.ResponseContent	""
// @Failure 50001 {object} UserController.ResponseContent	""
// @Router /logout [post]
func Logout(ctx iris.Context) {
	UserController.Logout(ctx)
}

// @Description
// @Accept  json
// @Produce  json
// @Param param body UserController.UserRegistrationReq true "param"
// @Success 20000 {object} UserController.ResponseContent	""
// @Failure 50001 {object} UserController.ResponseContent	""
// @Router /user/registration [post]
func Registration(ctx iris.Context) {
	UserController.Registration(ctx)
}