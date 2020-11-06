package register

import (
	"github.com/kataras/iris/v12"
	"github.com/hfeng101/niwo/apis/catalog/v1"
)

// 主题分类

func CatalogRouteRegister(app *iris.Application) {
	catalogParty := app.Party("/catalog")

	catalogParty.Get("listByKey",v1.GetRecordListByKey)
	catalogParty.Get("list",v1.GetRecordList)

	catalogParty.Get("personageRecordList",v1.GetPersonageRecordList)
	catalogParty.Get("sportRecordList",v1.GetSportRecordList)
	catalogParty.Get("economicsRecordList",v1.GetEconomicsRecordList)
	catalogParty.Get("militaryRecordList",v1.GetMilitaryRecordList)
	catalogParty.Get("entertainmentRecordList",v1.GetEntertainmentRecordList)
}
