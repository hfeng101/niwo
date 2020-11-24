package storage

import (
	AliOssSdk "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cihub/seelog"
	"github.com/hfeng101/niwo/utils/consts"
	"sync"
)

var (
	//CosHandle *TencentCosSdk.Client
	//CosBucketHandle *TencentCosSdk.Bucket
	//TencentSecretId string
	//TencentSecretKey string

	CosHandle *AliOssSdk.Client
	CosBucketHandle *AliOssSdk.Bucket
	AliSecretId string
	AliSecretKey string

	cosLock = sync.RWMutex{}
)

//腾讯云是有点烂，文档烂七八糟
//func init() {
//	bucketUrl,err := url.Parse("test123"+"")
//	if err != nil {
//		seelog.Errorf("Parse url failed, err is %v", err.Error())
//		return
//	}
//	baseUrl := &TencentCosSdk.BaseURL{BucketURL: bucketUrl}
//	CosHandle :=TencentCosSdk.NewClient(baseUrl, &http.Client{
//		Transport: &TencentCosSdk.AuthorizationTransport{
//			SecretID: TencentSecretId,
//			SecretKey: TencentSecretKey,
//		},
//	})
//
//	if CosHandle == nil {
//		seelog.Errorf("TencentCosSdk.NewClient failed, result is nil")
//		return
//	}
//}
//
//func buildBuckets(){
//	_,err := CosHandle.Bucket.Put(context.Background(), nil)
//}

//阿里云对象存储
func InitCos() {
	endpoint := "http://oss-cn-shenzhen.aliyuncs.com"

	CosHandle, err := AliOssSdk.New(endpoint, AliSecretId, AliSecretKey)
	if err != nil {
		seelog.Errorf("AliOssSdk.New failed, err is %v", err.Error())
		return
	}

	//为每一个分类创建一个桶
	if err := CosHandle.CreateBucket(consts.PERSONAGE);err != nil {
		seelog.Errorf("create bucket %v failed, err is %v", consts.PERSONAGE, err.Error())
		return
	}

	if err := CosHandle.CreateBucket(consts.SPORT);err != nil {
		seelog.Errorf("create bucket %v failed, err is %v", consts.SPORT, err.Error())
		return
	}

	if err := CosHandle.CreateBucket(consts.SPORT);err != nil {
		seelog.Errorf("create bucket %v failed, err is %v", consts.SPORT, err.Error())
		return
	}

	if err := CosHandle.CreateBucket(consts.ECONOMICS);err != nil {
		seelog.Errorf("create bucket %v failed, err is %v", consts.SPORT, err.Error())
		return
	}

	if err := CosHandle.CreateBucket(consts.MILITARY);err != nil {
		seelog.Errorf("create bucket %v failed, err is %v", consts.SPORT, err.Error())
		return
	}

	if err := CosHandle.CreateBucket(consts.ENTERTAINMENT);err != nil {
		seelog.Errorf("create bucket %v failed, err is %v", consts.SPORT, err.Error())
		return
	}
}

func GetCosHandle() interface{}{
	cosLock.RLock()
	defer cosLock.RUnlock()

	return CosHandle
}