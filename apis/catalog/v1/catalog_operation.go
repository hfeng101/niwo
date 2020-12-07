package v1

import (
	"github.com/kataras/iris/v12"
	CatalogController "github.com/hfeng101/niwo/controller/catalog/v1"
	//"github.com/hfeng101/niwo"
	//operation_controll "github.com/hfeng101/niwo/controller/catalog/v1"
)

// @Title catalog operation API
// @Version 1.0
// @Contact.name catalog operation
// @Contact.url http://www.niwofly.com
// @Description 根据搜索获取关键字对应的主题列表
// @Accept  json
// @Produce  json
// @Param param body CatalogController.GetRecordListByKeywordReq true "param"
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /v1/catalog/listByKey [post]
// @Host niwofly.com
// @BasePath /v1
func GetRecordListByKey(ctx iris.Context) {
	CatalogController.GetRecordListByKey(ctx)
}

// @Title catalog operation API
// @Version 1.0
// @Contact.name catalog operation
// @Contact.url http://www.niwofly.com
// @Description 获取指定分类的纪录主题列表
// @Accept  json
// @Produce  json
// @Param catalog query string true "param"
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /v1/catalog/list [get]
// @Host niwofly.com
// @BasePath /v1
func GetRecordList(ctx iris.Context) {
	CatalogController.GetRecordList(ctx)
}

// @Title catalog operation API
// @Version 1.0
// @Contact.name catalog operation
// @Contact.url http://www.niwofly.com
// @Description 获取"人物"分类的纪录主题列表
// @Accept  json
// @Produce  json
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /v1/catalog/personageRecordList [get]
// @Host niwofly.com
// @BasePath /v1
func GetPersonageRecordList(ctx iris.Context) {
	CatalogController.GetPersonageRecordList(ctx)
}

// @Title catalog operation API
// @Version 1.0
// @Contact.name catalog operation
// @Contact.url http://www.niwofly.com
// @Description 获取"体育"分类的纪录主题列表
// @Accept  json
// @Produce  json
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /v1/catalog/sportRecordList [get]
// @Host niwofly.com
// @BasePath /v1
func GetSportRecordList(ctx iris.Context) {
	CatalogController.GetSportRecordList(ctx)
}

// @Title catalog operation API
// @Version 1.0
// @Contact.name catalog operation
// @Contact.url http://www.niwofly.com
// @Description 获取"经济"分类的纪录主题列表
// @Accept  json
// @Produce  json
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /v1/catalog/economicsRecordList [get]
// @Host niwofly.com
// @BasePath /v1
func GetEconomicsRecordList(ctx iris.Context) {
	CatalogController.GetEconomicsRecordList(ctx)
}

// @Title catalog operation API
// @Version 1.0
// @Contact.name catalog operation
// @Contact.url http://www.niwofly.com
// @Description 获取"军事"分类的纪录主题列表
// @Accept  json
// @Produce  json
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /v1/catalog/militaryRecordList [get]
// @Host niwofly.com
// @BasePath /v1
func GetMilitaryRecordList(ctx iris.Context) {
	CatalogController.GetMilitaryRecordList(ctx)
}

// @Title catalog operation API
// @Version 1.0
// @Contact.name catalog operation
// @Contact.url http://www.niwofly.com
// @Description 获取"娱乐"分类的纪录主题列表
// @Accept  json
// @Produce  json
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /v1/catalog/entertainmentRecordList [get]
// @Host niwofly.com
// @BasePath /v1
func GetEntertainmentRecordList(ctx iris.Context) {
	CatalogController.GetEntertainmentRecordList(ctx)
}

