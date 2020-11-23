package v1

import (
	"github.com/cihub/seelog"
	"github.com/hfeng101/niwo/storage"
	"github.com/hfeng101/niwo/utils/consts"
	"github.com/kataras/iris/v12"
)

//根据关键字获取记录列表
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func GetRecordListByKey(ctx iris.Context) {
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &GetRecordListByKeywordReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	switch req.catalog {
		case consts.PERSONAGE:
			personageRecordList,err := GetListByKeywordFromPersonageRecord(req.Keyword)
			if err!= nil {
				seelog.Errorf("GetListByKeywordFromPersonageRecord failed ,err is %v", err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
			resBody.Data = personageRecordList
			return
		case consts.SPORT:
			sportRecordList,err := GetListByKeywordFromSportRecord(req.Keyword)
			if err != nil {
				seelog.Errorf("GetListByKeywordFromSportRecord failed ,err is %v", err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
			resBody.Data = sportRecordList
			return
		case consts.ECONOMICS:
			economicsRecordList,err := GetListByKeywordFromEconomicsRecord(req.Keyword)
			if err != nil {
				seelog.Errorf("GetListByKeywordFromEconomicsRecord failed ,err is %v", err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
			resBody.Data = economicsRecordList
			return
		case consts.MILITARY:
			militaryRecordList,err := GetListByKeywordFromMilitaryRecord(req.Keyword)
			if err != nil {
				seelog.Errorf("GetListByKeywordFromMilitaryRecord failed ,err is %v", err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
			resBody.Data = militaryRecordList
			return
		case consts.ENTERTAINMENT:
			entertainmentRecordList,err := GetListByKeywordFromEntertainmentRecord(req.Keyword)
			if err != nil {
				seelog.Errorf("GetListByKeywordFromEntertainmentRecord failed ,err is %v", err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
			resBody.Data = entertainmentRecordList
			return
			//ctx.JSON(entertainmentRecordList)
		default:
			recordList,err := GetListByKeywordFromAllRecord(req.Keyword)
			if err != nil {
				seelog.Errorf("GetListByKeywordFromAllRecord failed ,err is %v", err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
			resBody.Data = recordList
			return
	}

}

//按catalog匹配查看记录列表
func GetRecordList(ctx iris.Context) {
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &GetRecordListReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	//从数据库里获取列表，并分列显示
	switch req.catalog {
		case consts.PERSONAGE:
			personageRecordList := []*storage.PersonageRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := storage.GetMysqlDbHandle()
			//关键字模糊查找
			if dbHandle.Find(personageRecordList) == nil {
				seelog.Errorf("Get personage record list failed from mysql failed")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			resBody.Data = personageRecordList
		case consts.SPORT:
			sportRecordList := []*storage.SportRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := storage.GetMysqlDbHandle()
			//关键字模糊查找
			if dbHandle.Find(sportRecordList) == nil {
				seelog.Errorf("Get sport record list failed from mysql failed")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			resBody.Data = sportRecordList
		case consts.ECONOMICS:
			economicsRecordList := []*storage.EconomicsRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := storage.GetMysqlDbHandle()
			//关键字模糊查找
			if dbHandle.Find(economicsRecordList) == nil {
				seelog.Errorf("Get economics record list failed from mysql failed")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			resBody.Data = economicsRecordList
		case consts.MILITARY:
			militaryRecordList := []*storage.MilitaryRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := storage.GetMysqlDbHandle()
			//关键字模糊查找
			if dbHandle.Find(militaryRecordList) == nil {
				seelog.Errorf("Get military record list failed from mysql failed")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			resBody.Data = militaryRecordList
		case consts.ENTERTAINMENT:
			entertainmentRecordList := []*storage.EntertainmentRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := storage.GetMysqlDbHandle()
			//关键字模糊查找
			if dbHandle.Find(entertainmentRecordList) == nil {
				seelog.Errorf("Get entertainment record list failed from mysql failed")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			resBody.Data = entertainmentRecordList
		default:
			seelog.Errorf("Catalog type is invalid, cannot found catalog")
			resBody.Code = consts.ERRORCODE
			resBody.Message = consts.ERRORCODEMESSAGE + "Catalog type is invalid, cannot find"
			return
	}
}

//人物记录列表
func GetPersonageRecordList(ctx iris.Context) {
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &GetPersonageRecordListReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	personageRecordList := []*storage.PersonageRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	if dbHandle.Find(personageRecordList) == nil {
		seelog.Errorf("Get personage record list failed from mysql failed")
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE
		return
	}

	resBody.Data = personageRecordList

	return
}

//体育记录列表
func GetSportRecordList(ctx iris.Context) {
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &GetSportRecordListReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	sportRecordList := []*storage.SportRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	if dbHandle.Find(sportRecordList) == nil {
		seelog.Errorf("Get sport record list failed from mysql failed")
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE
		return
	}

	resBody.Data = sportRecordList
}

//经济记录列表
func GetEconomicsRecordList(ctx iris.Context) {
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &GetEconomicsRecordListReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	economicsRecordList := []*storage.EconomicsRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	if dbHandle.Find(economicsRecordList) == nil {
		seelog.Errorf("Get economics record list failed from mysql failed")
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE
		return
	}

	resBody.Data = economicsRecordList
}

//军事记录列表
func GetMilitaryRecordList(ctx iris.Context) {
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &GetMilitaryRecordListReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	militaryRecordList := []*storage.MilitaryRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	if dbHandle.Find(militaryRecordList) == nil {
		seelog.Errorf("Get military record list failed from mysql failed")
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE
		return
	}

	resBody.Data = militaryRecordList
}

//娱乐记录列表
func GetEntertainmentRecordList(ctx iris.Context) {
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &GetEntertainmentRecordListReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	entertainmentRecordList := []*storage.EntertainmentRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	if dbHandle.Find(entertainmentRecordList) == nil {
		seelog.Errorf("Get entertainment record list failed from mysql failed")
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE
		return
	}

	resBody.Data = entertainmentRecordList
}