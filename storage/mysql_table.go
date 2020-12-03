package storage

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
	HeadImage string `gorm:"column:head_image;type:varchar(256)"`
}

func (*UserInfo)TableName() string{
	return "user_info"
}

//用户视图，收藏、浏览记录、粉丝、关注、话题数据
type UserView struct {
	gorm.Model
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`
	Name string `gorm:"column:name;type:varchar(128)"`
	CollectionNumber string `gorm:"column:collection_number;type:int"`
	HistoryNumber string `gorm:"column:history_number;type:int"`
	FansNumber string `gorm:"column:fans_number;type:int"`
	FollowNumber string `gorm:"column:follow_number;type:int"`
	TopicNumber string `gorm:"column:topic_number;type:int"`
}

func (*UserView)TableName() string{
	return "user_view"
}

//收藏列表
type CollectionList struct {
	gorm.Model
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`
	Name string `gorm:"column:name;type:varchar(128)"`
}

func (*CollectionList)TableName() string{
	return "collection_list"
}

//浏览记录列表
type HistoryList struct {
	gorm.Model
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`
	Name string `gorm:"column:name;type:varchar(128)"`
}

func (*HistoryList)TableName() string{
	return "history_list"
}

//关注列表
type FollowList struct {
	gorm.Model
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`
	Name string `gorm:"column:name;type:varchar(128)"`
}

func (*FollowList)TableName() string{
	return "follow_list"
}

//粉丝列表
type FansList struct {
	gorm.Model
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`
	Name string `gorm:"column:name;type:varchar(128)"`
}

func (*FansList)TableName() string{
	return "fans_list"
}

//话题列表
type TopicList struct {
	gorm.Model
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`
	Name string `gorm:"column:name;type:varchar(128)"`
}

func (*TopicList)TableName() string{
	return "topic_list"
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

//上传文件（普通文件、图片、语音、长视频，短视频)的对象存储元数据(objectKey）记录表
type OsObjectInfoList struct {
	gorm.Model
	Catalog string `json:"colume:catalog;type:varchar(32)"`
	ObjectKey string `json:"colume:object_key;type:varchar(128)"`
}

func (*OsObjectInfoList)TableName() string{
	return "os_object_info_list"
}