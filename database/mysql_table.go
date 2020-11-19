package database

import "github.com/jinzhu/gorm"

//用户信息
type UserInfo struct {
	gorm.Model
	Uuid string `gorm:"column:uuid;type:varchar(128)"`
	Name string `gorm:"column:name;type:varchar(128)"`
	Password string `gorm:"column:password;type:varchar(128)"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`
	VerificationCode string `gorm:"column:verification_code;type:varchar(8)"`
	Email string `gorm:"column:email;type:varchar(128)"`
	Token string `gorm:"column:token;type:varchar(64)"`
	RefreshToken string `gorm:"column:refresh_token;type:varchar(64)"`
	Secret string `gorm:"column:secret;type:varchar(8)"`
}

func (*UserInfo)TableName() string{
	return "user_info"
}

//主题分类
type ThemeCatalog struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(128)"`
}

func (*ThemeCatalog)TableName() string{
	return "theme_catalog"
}

//通用记录列表数据结构
type GenerationRecord struct {
	gorm.Model
	Uuid string `json:"colume:uuid;type:varchar(128)"`
	Theme string `json:"colume:theme;type:varchar(128)"`
	Keyword string `json:"colume:keyword;type:varchar(128)"`

	//mongodb内容存放的键值索引
}

//人物记录
type PersonageRecordList struct {
	gorm.Model
	Uuid string `json:"colume:uuid;type:varchar(128)"`
	Theme string `json:"colume:theme;type:varchar(128)"`
	Keyword string `json:"colume:keyword;type:varchar(128)"`
	//map[string]interface{}
}

func (*PersonageRecordList)TableName() string{
	return "personage_record_list"
}

//经济记录
type EconomicsRecordList struct {
	gorm.Model
	Uuid string `json:"colume:uuid;type:varchar(128)"`
	Theme string `json:"colume:Theme;type:varchar(128)"`
	Keyword string `json:"colume:keyword;type:varchar(128)"`
}

func (*EconomicsRecordList)TableName() string{
	return "economics_record_list"
}

//体育记录
type SportRecordList struct {
	gorm.Model
	Uuid string `json:"colume:uuid;type:varchar(128)"`
	Theme string `json:"colume:Theme;type:varchar(128)"`
	Keyword string `json:"colume:keyword;type:varchar(128)"`
}

func (*SportRecordList)TableName() string{
	return "sport_record_list"
}

//军事记录
type MilitaryRecordList struct {
	gorm.Model
	Uuid string `json:"colume:uuid;type:varchar(128)"`
	Theme string `json:"colume:Theme;type:varchar(128)"`
	Keyword string `json:"colume:keyword;type:varchar(128)"`
}

func (*MilitaryRecordList)TableName() string{
	return "military_record_list"
}

//娱乐记录
type EntertainmentRecordList struct {
	gorm.Model
	Uuid string `json:"colume:uuid;type:varchar(128)"`
	Theme string `json:"colume:Theme;type:varchar(128)"`
	Keyword string `json:"colume:keyword;type:varchar(128)"`
}

func (*EntertainmentRecordList)TableName() string{
	return "entertainment_record_list"
}