package mysql

import (
	"github.com/cihub/seelog"
	"github.com/hfeng101/niwo/utils/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"time"
)

var (
	DbHandle = &gorm.DB{}
	MysqlLock = sync.RWMutex{}
)

func InitGlobalOrm() error{
	var err error

	//DbHandle,err:=sql.Open("mysql",userName+":"+password+"@tcp("+ip+")/"+dbName+"?charset=utf8")
	//全局变量和局部变量搞混，赋值就近原则，导致全局变量为正常赋值，所以注释掉
	//DbHandle,err := gorm.Open("mysql", config.GetConfig().Other["mysql-dev"])
	DbHandle,err = gorm.Open("mysql", config.GetConfig().Other["mysql-dev"])
	if err != nil {
		seelog.Errorf("Open mysql failed, err is %v",err.Error())
		return err
	}

	DbHandle.DB().SetConnMaxLifetime(time.Duration(config.GetConfig().Other["mysql-maxlifetime"].(int)) * time.Second)
	DbHandle.DB().SetMaxIdleConns(config.GetConfig().Other["mysql-max-idle"].(int))
	DbHandle.DB().SetMaxOpenConns(config.GetConfig().Other["mysql-max-open"].(int))
	DbHandle.LogMode(config.GetConfig().Other["mysql-log-mode"].(bool))

	return nil
}

func GetMysqlDbHandle() *gorm.DB {
	MysqlLock.RLock()
	MysqlLock.RUnlock()
	return DbHandle
}

