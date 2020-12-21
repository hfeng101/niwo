package cloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/cihub/seelog"
	"github.com/hfeng101/niwo/utils/consts"
	"sync"
)

var (
	DysmsapiHandle *dysmsapi.Client
	Lock = &sync.RWMutex{}
)

func InitAliyun() error{
	var err error
	DysmsapiHandle, err = dysmsapi.NewClientWithAccessKey("cn-shenzhen", consts.AliSecretId, consts.AliSecretKey)
	if err != nil {
		seelog.Errorf("dysmsapi NewClientWithAccessKey failed, err is %v", err.Error())
		return err
	}

	return nil
}

func GetDysmsapiHandle() *dysmsapi.Client {
	Lock.RLock()
	defer Lock.RUnlock()

	return DysmsapiHandle
}

func DestroyDysmsapiHandle() {
	Lock.RLock()
	defer Lock.RUnlock()

	//DysmsapiHandle.
}