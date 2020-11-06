package database

import (
	"sync"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/cihub/seelog"
	"github.com/hfeng101/niwo/utils/config"
)

var (
	DbHandle *gorm.DB

	MysqlLock = sync.RWMutex{}
)

func InitGlobalOrm() {
	//DbHandle,err:=sql.Open("mysql",userName+":"+password+"@tcp("+ip+")/"+dbName+"?charset=utf8")
	DbHandle,err := gorm.Open("mysql", config.GetConfig().Other["mysql-dev"]);
	if  err != nil {
		DbHandle = nil
		seelog.Errorf("Open mysql failed, err is %v",err.Error())
		return
	}

	DbHandle.DB().SetConnMaxLifetime(time.Duration(config.GetConfig().Other["mysql-maxlifetime"].(int)) * time.Second)
	DbHandle.DB().SetMaxIdleConns(config.GetConfig().Other["mysql-max-idle"].(int))
	DbHandle.DB().SetMaxOpenConns(config.GetConfig().Other["mysql-max-open"].(int))
	DbHandle.LogMode(config.GetConfig().Other["mysql-log-mode"].(bool))
}

func GetMysqlDbHandle() *gorm.DB {
	MysqlLock.RLock()
	MysqlLock.RUnlock()

	return DbHandle
}

