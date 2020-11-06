package register

import "github.com/kataras/iris/v12"

//群组聊天窗口
func GroupRouteRegister(app *iris.Application) {
	groupParty := app.Party("/group")

	groupParty.Get("/list",)
}