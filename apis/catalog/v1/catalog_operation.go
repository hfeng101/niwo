package v1

import (
	"github.com/kataras/iris/v12"
	CatalogController "github.com/hfeng101/niwo/controller/catalog/v1"
	//"github.com/hfeng101/niwo"
	//operation_controll "github.com/hfeng101/niwo/controller/catalog/v1"
)

// @title catalog operation API
// @version 1.0
// @contact.name catalog operation
// @contact.url http://www.niwo.com
// @Description Get record list by keyword search.
// @Accept  json
// @Produce  json
// @Param param body CatalogController.GetRecordListByKeywordReq true "param"
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /catalog/listByKey [get]
// @host niwo.com
// @BasePath /v1
func GetRecordListByKey(ctx iris.Context) {
	CatalogController.GetRecordListByKey(ctx)
}

// @title catalog operation API
// @version 1.0
// @contact.name catalog operation
// @contact.url http://www.niwo.com
// @Description Get record list by keyword search.
// @Accept  json
// @Produce  json
// @Param param body CatalogController.GetRecordListReq true "param"
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /catalog/list [get]
// @host niwo.com
// @BasePath /v1
func GetRecordList(ctx iris.Context) {
	CatalogController.GetRecordList(ctx)
}

// @title catalog operation API
// @version 1.0
// @contact.name catalog operation
// @contact.url http://www.niwo.com
// @Description Get record list of personage record
// @Accept  json
// @Produce  json
// @Param param body CatalogController.GetPersonageRecordListReq true "param"
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /catalog/personageRecordList [get]
// @host niwo.com
// @BasePath /v1
func GetPersonageRecordList(ctx iris.Context) {
	CatalogController.GetPersonageRecordList(ctx)
}

// @title catalog operation API
// @version 1.0
// @contact.name catalog operation
// @contact.url http://www.niwo.com
// @Description Get record list of sport record
// @Accept  json
// @Produce  json
// @Param param body CatalogController.GetSportRecordListReq true "param"
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /catalog/sportRecordList [get]
// @host niwo.com
// @BasePath /v1
func GetSportRecordList(ctx iris.Context) {
	CatalogController.GetSportRecordList(ctx)
}

// @title catalog operation API
// @version 1.0
// @contact.name catalog operation
// @contact.url http://www.niwo.com
// @Description Get record list of economics record
// @Accept  json
// @Produce  json
// @Param param body CatalogController.GetEconomicsRecordListReq true "param"
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /catalog/economicsRecordList [get]
// @host niwo.com
// @BasePath /v1
func GetEconomicsRecordList(ctx iris.Context) {
	CatalogController.GetEconomicsRecordList(ctx)
}

// @title catalog operation API
// @version 1.0
// @contact.name catalog operation
// @contact.url http://www.niwo.com
// @Description Get record list of military record
// @Accept  json
// @Produce  json
// @Param param body CatalogController.GetMilitaryRecordListReq true "param"
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /catalog/militaryRecordList [get]
// @host niwo.com
// @BasePath /v1
func GetMilitaryRecordList(ctx iris.Context) {
	CatalogController.GetMilitaryRecordList(ctx)
}

// @title catalog operation API
// @version 1.0
// @contact.name catalog operation
// @contact.url http://www.niwo.com
// @Description Get record list of entertainment record
// @Accept  json
// @Produce  json
// @Param param body CatalogController.GetEntertainmentRecordListReq true "param"
// @Success 20000 {object} CatalogController.ResponseContent	""
// @Failure 50001 {object} CatalogController.ResponseContent	""
// @Router /catalog/entertainmentRecordList [get]
// @host niwo.com
// @BasePath /v1
func GetEntertainmentRecordList(ctx iris.Context) {
	CatalogController.GetEntertainmentRecordList(ctx)
}

