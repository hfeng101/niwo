package v1

import (
	"github.com/kataras/iris/v12"
	ContentController "github.com/hfeng101/niwo/controller/content/v1"
)

// @title Get Content API
// @version 1.0
// @contact.name Content operation of theme
// @contact.url http://www.niwo.com
// @Description Get content of a record
// @Accept  json
// @Produce  json
// @Param Param body ContentController.GetContentReq true "param"
// @Success 20000 {object} ContentController.ResponseContent	""
// @Failure 50001 {object} ContentController.ResponseContent	""
// @Router /content/theme [get]
// @host niwo.com
// @BasePath /v1
func GetContent(ctx iris.Context) {
	ContentController.GetContent(ctx)
}

// @title Add Content API
// @version 1.0
// @contact.name Content operation of theme
// @contact.url http://www.niwo.com
// @Description Add content of a record
// @Accept  json
// @Produce  json
// @Param Param body ContentController.AddContentReq true "param"
// @Success 20000 {object} ContentController.ResponseContent	""
// @Failure 50001 {object} ContentController.ResponseContent	""
// @Router /content/theme [post]
// @host niwo.com
// @BasePath /v1
func AddContent(ctx iris.Context) {
	ContentController.AddContent(ctx)
}

// @title Update Content API
// @version 1.0
// @contact.name Content operation of theme
// @contact.url http://www.niwo.com
// @Description Update content of a record
// @Accept  json
// @Produce  json
// @Param Param body ContentController.UpdateContentReq true "param"
// @Success 20000 {object} ContentController.ResponseContent	""
// @Failure 50001 {object} ContentController.ResponseContent	""
// @Router /content/theme [put]
// @host niwo.com
// @BasePath /v1
func UpdateContent(ctx iris.Context) {
	ContentController.UpdateContent(ctx)
}