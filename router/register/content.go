package register

import (
	"github.com/cihub/seelog"
	"github.com/kataras/iris/v12"
	"github.com/hfeng101/niwo/apis/content/v1"
)

//主要查看类容
func ContentRouteRegister(app *iris.Application) {
	contentParty := app.Party("/v1/content")

	seelog.Infof("ContentRouteRegister starting")
	//获取主题对应的内容
	contentParty.Get("/theme",v1.GetContent)

	//添加内容
	contentParty.Post("/theme",v1.AddContent)

	//更新内容
	contentParty.Put("/theme",v1.UpdateContent)

	//上传文件中引用的图片或视频
	contentParty.Post("/uploadFile",v1.UploadFile)

	//获取应用的图片或视频
	contentParty.Get("/getReferenceFile",v1.GetReferenceFile)
}
