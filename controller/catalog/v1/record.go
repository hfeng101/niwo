package v1

import (
	"github.com/hfeng101/niwo/database"
)

func GetListByKeywordFromPersonageRecord(key string)  ([]*database.PersonageRecordList, error){
	personageRecordList := []*database.PersonageRecordList{}

	dbHandle := database.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(personageRecordList)

	return personageRecordList,nil
}

func GetListByKeywordFromSportRecord(key string) ([]*database.SportRecordList,error) {
	sportRecordList := []*database.SportRecordList{}

	dbHandle := database.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(sportRecordList)

	return sportRecordList,nil
}

func GetListByKeywordFromEconomicsRecord(key string) ([]*database.EconomicsRecordList,error){
	economicsRecordList := []*database.EconomicsRecordList{}

	dbHandle := database.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(economicsRecordList)

	return economicsRecordList,nil
}

func GetListByKeywordFromMilitaryRecord(key string) ([]*database.MilitaryRecordList,error){
	militaryRecordList := []*database.MilitaryRecordList{}

	dbHandle := database.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(militaryRecordList)

	return militaryRecordList,nil
}

func GetListByKeywordFromEntertainmentRecord(key string) ([]*database.EntertainmentRecordList,error){
	entertainmentRecordList := []*database.EntertainmentRecordList{}

	dbHandle := database.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(entertainmentRecordList)

	return entertainmentRecordList,nil
}

//所有表合集
func GetListByKeywordFromAllRecord(key string) (interface{},error){

	return nil,nil
}