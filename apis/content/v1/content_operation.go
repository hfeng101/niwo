package v1

import (
	"github.com/kataras/iris/v12"
	ContentController "github.com/hfeng101/niwo/controller/content/v1"
)

// @title Get Content API
// @version 1.0
// @contact.name content operation
// @contact.url http://www.niwofly.com
// @Description 获取指定主题的具体内容
// @Accept  json
// @Produce  json
// @Param Param body ContentController.GetContentReq true "param"
// @Success 20000 {object} ContentController.ResponseContent	""
// @Failure 50001 {object} ContentController.ResponseContent	""
// @Router /v1/content/theme [get]
// @host niwofly.com
// @BasePath /v1
func GetContent(ctx iris.Context) {
	ContentController.GetContent(ctx)
}

// @title Add Content API
// @version 1.0
// @contact.name content operation
// @contact.url http://www.niwofly.com
// @Description 用户创作主题及对应内容
// @Accept  json
// @Produce  json
// @Param Param body ContentController.AddContentReq true "param"
// @Success 20000 {object} ContentController.ResponseContent	""
// @Failure 50001 {object} ContentController.ResponseContent	""
// @Router /v1/content/theme [post]
// @host niwofly.com
// @BasePath /v1
func AddContent(ctx iris.Context) {
	ContentController.AddContent(ctx)
}

// @title Update Content API
// @version 1.0
// @contact.name content operation
// @contact.url http://www.niwofly.com
// @Description 用户编辑更新内容
// @Accept  json
// @Produce  json
// @Param Param body ContentController.UpdateContentReq true "param"
// @Success 20000 {object} ContentController.ResponseContent	""
// @Failure 50001 {object} ContentController.ResponseContent	""
// @Router /v1/content/theme [put]
// @host niwofly.com
// @BasePath /v1
func UpdateContent(ctx iris.Context) {
	ContentController.UpdateContent(ctx)
}

// @title Update Content API
// @version 1.0
// @contact.name content operation
// @contact.url http://www.niwofly.com
// @Description 用户上传图片、音、视频文件
// @Accept  json
// @Produce  json
// @Param Param body ContentController.UpdateContentReq true "param"
// @Success 20000 {object} ContentController.ResponseContent	""
// @Failure 50001 {object} ContentController.ResponseContent	""
// @Router /v1/content/uploadFile [post]
// @host niwofly.com
// @BasePath /v1
func UploadFile(ctx iris.Context) {
	ContentController.UploadFile(ctx)
}

// @title Update Content API
// @version 1.0
// @contact.name content operation
// @contact.url http://www.niwofly.com
// @Description 用户获取图片、音、视频文件
// @Accept  json
// @Produce  json
// @Param Param body ContentController.GetReferenceFileReq true "param"
// @Success 20000 {object} ContentController.ResponseContent	""
// @Failure 50001 {object} ContentController.ResponseContent	""
// @Router /v1/content/getReferenceFile [get]
// @host niwofly.com
// @BasePath /v1
func GetReferenceFile(ctx iris.Context) {
	ContentController.GetReferenceFile(ctx)
}