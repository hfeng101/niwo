package v1

import (
	"github.com/cihub/seelog"
	"github.com/hfeng101/niwo/storage/mysql"
	"github.com/hfeng101/niwo/utils/consts"
	"github.com/kataras/iris/v12"
)

//根据关键字获取记录列表
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

	switch req.Catalog {
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
	defer ctx.JSON(resBody)

	catalog := ctx.Request().URL.Query().Get("catalog")

	//catalog = ctx.URLParam("catalog")
	//seelog.Infof("catalog 444444444 from url is %v", catalog)
	//从数据库里获取列表，并分列显示
	switch catalog {
		case consts.PERSONAGE:
			personageRecordList := &[]mysql.PersonageRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := mysql.GetMysqlDbHandle()
			//关键字模糊查找
			if dbHandle.Find(personageRecordList) == nil {
				seelog.Errorf("Get personage record list failed from mysql failed")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			resBody.Data = personageRecordList
		case consts.SPORT:
			sportRecordList := &[]mysql.SportRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := mysql.GetMysqlDbHandle()
			//关键字模糊查找
			if dbHandle.Find(sportRecordList) == nil {
				seelog.Errorf("Get sport record list failed from mysql failed")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			resBody.Data = sportRecordList
		case consts.ECONOMICS:
			economicsRecordList := &[]mysql.EconomicsRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := mysql.GetMysqlDbHandle()
			//关键字模糊查找
			if dbHandle.Find(economicsRecordList) == nil {
				seelog.Errorf("Get economics record list failed from mysql failed")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			resBody.Data = economicsRecordList
		case consts.MILITARY:
			militaryRecordList := &[]mysql.MilitaryRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := mysql.GetMysqlDbHandle()
			//关键字模糊查找
			if dbHandle.Find(militaryRecordList) == nil {
				seelog.Errorf("Get military record list failed from mysql failed")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			resBody.Data = militaryRecordList
		case consts.ENTERTAINMENT:
			entertainmentRecordList := &[]mysql.EntertainmentRecordList{}
			//从数据库里获取列表，并分列显示
			dbHandle := mysql.GetMysqlDbHandle()
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
	//req := &GetPersonageRecordListReq{}

	defer ctx.JSON(resBody)
	//if err := ctx.ReadJSON(req); err != nil {
	//	seelog.Errorf("Input param is valid, err is %v",err.Error())
	//	resBody.Code = consts.ERRORCODE
	//	resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
	//	return
	//}

	personageRecordList := &[]mysql.PersonageRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := mysql.GetMysqlDbHandle()
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
	//req := &GetSportRecordListReq{}

	defer ctx.JSON(resBody)
	//if err := ctx.ReadJSON(req); err != nil {
	//	seelog.Errorf("Input param is valid, err is %v",err.Error())
	//	resBody.Code = consts.ERRORCODE
	//	resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
	//	return
	//}

	sportRecordList := &[]mysql.SportRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := mysql.GetMysqlDbHandle()
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
	//req := &GetEconomicsRecordListReq{}

	defer ctx.JSON(resBody)
	//if err := ctx.ReadJSON(req); err != nil {
	//	seelog.Errorf("Input param is valid, err is %v",err.Error())
	//	resBody.Code = consts.ERRORCODE
	//	resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
	//	return
	//}

	economicsRecordList := &[]mysql.EconomicsRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := mysql.GetMysqlDbHandle()
	//关键字模糊查找
	if dbHandle.Where("").Find(economicsRecordList) == nil {
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
	//req := &GetMilitaryRecordListReq{}

	defer ctx.JSON(resBody)
	//if err := ctx.ReadJSON(req); err != nil {
	//	seelog.Errorf("Input param is valid, err is %v",err.Error())
	//	resBody.Code = consts.ERRORCODE
	//	resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
	//	return
	//}

	militaryRecordList := &[]mysql.MilitaryRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := mysql.GetMysqlDbHandle()
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
	//req := &GetEntertainmentRecordListReq{}

	defer ctx.JSON(resBody)
	//if err := ctx.ReadJSON(req); err != nil {
	//	seelog.Errorf("Input param is valid, err is %v",err.Error())
	//	resBody.Code = consts.ERRORCODE
	//	resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
	//	return
	//}

	entertainmentRecordList := &[]mysql.EntertainmentRecordList{}
	//从数据库里获取列表，并分列显示
	dbHandle := mysql.GetMysqlDbHandle()
	//关键字模糊查找
	if dbHandle.Find(entertainmentRecordList) == nil {
		seelog.Errorf("Get entertainment record list failed from mysql failed")
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE
		return
	}

	resBody.Data = entertainmentRecordList
}