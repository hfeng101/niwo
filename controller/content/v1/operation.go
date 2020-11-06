package v1

import (
	"github.com/cihub/seelog"
	uuid2 "github.com/google/uuid"
	"github.com/hfeng101/niwo/database"
	"github.com/hfeng101/niwo/utils/consts"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

func GetContent(ctx iris.Context){
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &GetContentReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	//联合查询，从mysql查询到mongodb的键值，然后去mongodb查内容
	//record := &database.GenerationRecord{}
	//dbHandle := database.GetMysqlDbHandle()
	//dbHandle.Where("uuid = ", req.Uuid).Find(record)

	//从mongodb获取对应UUID中的内容
	content := &database.GenerationRecordContent{}
	mongoDbHandle := database.GetMongoDbHandle()

	filter := bson.M{
		"uuid": req.Uuid,
	}

	switch req.CatalogType {
		case consts.PERSONAGE:
			mongoDbHandle.Collection(consts.PERSONAGE).FindOne(ctx.Request().Context(),filter).Decode(content)
		case consts.SPORT:
			mongoDbHandle.Collection(consts.SPORT).FindOne(ctx.Request().Context(),filter).Decode(content)
		case consts.ECONOMICS:
			mongoDbHandle.Collection(consts.ECONOMICS).FindOne(ctx.Request().Context(),filter).Decode(content)
		case consts.MILITARY:
			mongoDbHandle.Collection(consts.MILITARY).FindOne(ctx.Request().Context(),filter).Decode(content)
		case consts.ENTERTAINMENT:
			mongoDbHandle.Collection(consts.ENTERTAINMENT).FindOne(ctx.Request().Context(),filter).Decode(content)
		default :
			seelog.Errorf("Cannot find catalog type")
			resBody.Code = consts.ERRORCODE
			resBody.Message = consts.ERRORCODEMESSAGE + "Catalog type is invalid, cannot find"
			return
	}

	resBody.Data = content
	return
}

func AddContent(ctx iris.Context){
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &AddContentReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	//先生成uuid
	uuid := uuid2.New()
	mysqlDbHandle := database.GetMysqlDbHandle()

	//生成要写入mongodb的内容
	data := bson.M{
		"uuid": uuid,
		"content": req.Content,
	}
	mongoDbHandle := database.GetMongoDbHandle()
	switch req.CatalogType {
		case consts.PERSONAGE:
			//先把索引存入mysql
			count := 0
			peronageRecord := &database.PersonageRecordList{
				Uuid: uuid.String(),
				Theme: req.Theme,
				Keyword: req.Keyword,
			}
			mysqlDbHandle.FirstOrCreate(peronageRecord)
			//确认落库成功
			mysqlDbHandle.Where("").Find(peronageRecord).Count(&count)
			if count == 0 {
				seelog.Errorf("Mysql first create failed!!!")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			//将内容写入mongodb
			if _,err := mongoDbHandle.Collection(consts.PERSONAGE).InsertOne(ctx.Request().Context(), data);err != nil {
				seelog.Errorf("Insert content to collection %v failed, err is %v!!!", consts.PERSONAGE, err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
		case consts.SPORT:
			//先把索引存入mysql
			count := 0
			sportRecord := &database.SportRecordList{
				Uuid: uuid.String(),
				Theme: req.Theme,
				Keyword: req.Keyword,
			}
			mysqlDbHandle.FirstOrCreate(sportRecord)
			//确认落库成功
			mysqlDbHandle.Where("").Find(sportRecord).Count(&count)
			if count == 0 {
				seelog.Errorf("Mysql first create failed!!!")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			//将内容写入mongodb
			if _,err := mongoDbHandle.Collection(consts.SPORT).InsertOne(ctx.Request().Context(), data);err != nil {
				seelog.Errorf("Insert content to collection %v failed, err is %v!!!", consts.PERSONAGE, err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
		case consts.ECONOMICS:
			//先把索引存入mysql
			count := 0
			economicsRecord := &database.SportRecordList{
				Uuid: uuid.String(),
				Theme: req.Theme,
				Keyword: req.Keyword,
			}
			mysqlDbHandle.FirstOrCreate(economicsRecord)
			//确认落库成功
			mysqlDbHandle.Where("").Find(economicsRecord).Count(&count)
			if count == 0 {
				seelog.Errorf("Mysql first create failed!!!")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			//将内容写入mongodb
			if _,err := mongoDbHandle.Collection(consts.ECONOMICS).InsertOne(ctx.Request().Context(), data);err != nil {
				seelog.Errorf("Insert content to collection %v failed, err is %v!!!", consts.ECONOMICS, err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
		case consts.MILITARY:
			//先把索引存入mysql
			count := 0
			militaryRecord := &database.SportRecordList{
				Uuid: uuid.String(),
				Theme: req.Theme,
				Keyword: req.Keyword,
			}
			mysqlDbHandle.FirstOrCreate(militaryRecord)
			//确认落库成功
			mysqlDbHandle.Where("").Find(militaryRecord).Count(&count)
			if count == 0 {
				seelog.Errorf("Mysql first create failed!!!")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			//将内容写入mongodb
			if _,err := mongoDbHandle.Collection(consts.MILITARY).InsertOne(ctx.Request().Context(), data);err != nil {
				seelog.Errorf("Insert content to collection %v failed, err is %v!!!", consts.MILITARY, err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
		case consts.ENTERTAINMENT:
			//先把索引存入mysql
			count := 0
			sportRecord := &database.SportRecordList{
				Uuid: uuid.String(),
				Theme: req.Theme,
				Keyword: req.Keyword,
			}
			mysqlDbHandle.FirstOrCreate(sportRecord)
			//确认落库成功
			mysqlDbHandle.Where("").Find(sportRecord).Count(&count)
			//database.MysqlGlobalOrm.FirstOrCreate(kubeConf, `cfg_content = "`+kubeConf.CfgConetent+`"`)
			//database.MysqlGlobalOrm.Where("cfg_content = ?", kubeConf.CfgConetent).Find(&[]*v1.KubeConfig{}).Count(&count)
			if count == 0 {
				seelog.Errorf("Mysql first create failed!!!")
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}

			//将内容写入mongodb
			if _,err := mongoDbHandle.Collection(consts.ENTERTAINMENT).InsertOne(ctx.Request().Context(), data);err != nil {
				seelog.Errorf("Insert content to collection %v failed, err is %v!!!", consts.ENTERTAINMENT, err.Error())
				resBody.Code = consts.ERRORCODE
				resBody.Message = consts.ERRORCODEMESSAGE
				return
			}
		default:
			seelog.Errorf("Cannot find catalog type")
			resBody.Code = consts.ERRORCODE
			resBody.Message = consts.ERRORCODEMESSAGE + "Catalog type is invalid, cannot find"
			return
	}

	//mongoDbHandle.Collection()


}

func UpdateContent(ctx iris.Context){
	resBody := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
		Data: nil,
	}
	req := &UpdateContentReq{}

	defer ctx.JSON(resBody)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Input param is valid, err is %v",err.Error())
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	//将内容写入mongodb
	filter := bson.M{
		"uuid": req.Uuid,
	}

	data := bson.M{
		"uuid": req.Uuid,
		"content": req.Content,
	}
	mongoDbHandle := database.GetMongoDbHandle()
	switch req.CatalogType {
	case consts.PERSONAGE:
		if _,err := mongoDbHandle.Collection(consts.PERSONAGE).UpdateOne(ctx.Request().Context(), filter, data);err != nil {
			seelog.Errorf("Update content to collection %v failed, err is %v!!!", consts.PERSONAGE, err.Error())
			resBody.Code = consts.ERRORCODE
			resBody.Message = consts.ERRORCODEMESSAGE
			return
		}
	case consts.SPORT:
		if _,err := mongoDbHandle.Collection(consts.SPORT).UpdateOne(ctx.Request().Context(), filter, data);err != nil {
			seelog.Errorf("Update content to collection %v failed, err is %v!!!", consts.SPORT, err.Error())
			resBody.Code = consts.ERRORCODE
			resBody.Message = consts.ERRORCODEMESSAGE
			return
		}
	case consts.ECONOMICS:
		if _,err := mongoDbHandle.Collection(consts.ECONOMICS).UpdateOne(ctx.Request().Context(), filter, data);err != nil {
			seelog.Errorf("Update content to collection %v failed, err is %v!!!", consts.ECONOMICS, err.Error())
			resBody.Code = consts.ERRORCODE
			resBody.Message = consts.ERRORCODEMESSAGE
			return
		}
	case consts.MILITARY:
		if _,err := mongoDbHandle.Collection(consts.MILITARY).UpdateOne(ctx.Request().Context(), filter, data);err != nil {
			seelog.Errorf("Update content to collection %v failed, err is %v!!!", consts.MILITARY, err.Error())
			resBody.Code = consts.ERRORCODE
			resBody.Message = consts.ERRORCODEMESSAGE
			return
		}
	case consts.ENTERTAINMENT:
		if _,err := mongoDbHandle.Collection(consts.ENTERTAINMENT).UpdateOne(ctx.Request().Context(), filter, data);err != nil {
			seelog.Errorf("Update content to collection %v failed, err is %v!!!", consts.ENTERTAINMENT, err.Error())
			resBody.Code = consts.ERRORCODE
			resBody.Message = consts.ERRORCODEMESSAGE
			return
		}
	default:
		seelog.Errorf("Cannot find catalog type")
		resBody.Code = consts.ERRORCODE
		resBody.Message = consts.ERRORCODEMESSAGE + "Catalog type is invalid, cannot find"
		return
	}
}

