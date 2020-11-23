package v1

import (
	"github.com/hfeng101/niwo/storage"
)

func GetListByKeywordFromPersonageRecord(key string)  ([]*storage.PersonageRecordList, error){
	personageRecordList := []*storage.PersonageRecordList{}

	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(personageRecordList)

	return personageRecordList,nil
}

func GetListByKeywordFromSportRecord(key string) ([]*storage.SportRecordList,error) {
	sportRecordList := []*storage.SportRecordList{}

	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(sportRecordList)

	return sportRecordList,nil
}

func GetListByKeywordFromEconomicsRecord(key string) ([]*storage.EconomicsRecordList,error){
	economicsRecordList := []*storage.EconomicsRecordList{}

	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(economicsRecordList)

	return economicsRecordList,nil
}

func GetListByKeywordFromMilitaryRecord(key string) ([]*storage.MilitaryRecordList,error){
	militaryRecordList := []*storage.MilitaryRecordList{}

	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(militaryRecordList)

	return militaryRecordList,nil
}

func GetListByKeywordFromEntertainmentRecord(key string) ([]*storage.EntertainmentRecordList,error){
	entertainmentRecordList := []*storage.EntertainmentRecordList{}

	dbHandle := storage.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(entertainmentRecordList)

	return entertainmentRecordList,nil
}

//所有表合集
func GetListByKeywordFromAllRecord(key string) (interface{},error){

	return nil,nil
}