package v1

import (
	"github.com/hfeng101/niwo/storage/mysql"
)

func GetListByKeywordFromPersonageRecord(key string)  (*[]mysql.PersonageRecordList, error){
	personageRecordList := &[]mysql.PersonageRecordList{}

	dbHandle := mysql.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(personageRecordList)

	return personageRecordList,nil
}

func GetListByKeywordFromSportRecord(key string) (*[]mysql.SportRecordList,error) {
	sportRecordList := &[]mysql.SportRecordList{}

	dbHandle := mysql.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(sportRecordList)

	return sportRecordList,nil
}

func GetListByKeywordFromEconomicsRecord(key string) (*[]mysql.EconomicsRecordList,error){
	economicsRecordList := &[]mysql.EconomicsRecordList{}

	dbHandle := mysql.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(economicsRecordList)

	return economicsRecordList,nil
}

func GetListByKeywordFromMilitaryRecord(key string) (*[]mysql.MilitaryRecordList,error){
	militaryRecordList := &[]mysql.MilitaryRecordList{}

	dbHandle := mysql.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(militaryRecordList)

	return militaryRecordList,nil
}

func GetListByKeywordFromEntertainmentRecord(key string) (*[]mysql.EntertainmentRecordList,error){
	entertainmentRecordList := &[]mysql.EntertainmentRecordList{}

	dbHandle := mysql.GetMysqlDbHandle()
	//关键字模糊查找
	dbHandle.Where("theme like", "%"+key+"%").Find(entertainmentRecordList)

	return entertainmentRecordList,nil
}

//所有表合集
func GetListByKeywordFromAllRecord(key string) (interface{},error){

	return nil,nil
}