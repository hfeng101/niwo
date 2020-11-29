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
	AliSecretId = "LTAI4G8vW2yz1i71mDBS53hg"
	AliSecretKey = "rsiXMWlolkYZ3d5u3bZU5PJIFtV8Os"

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
func InitCos() error {
//func init() {
	var err error
	endpoint := "http://oss-cn-shenzhen.aliyuncs.com"

	CosHandle, err = AliOssSdk.New(endpoint, AliSecretId, AliSecretKey)
	if err != nil {
		seelog.Errorf("AliOssSdk.New failed, err is %v", err.Error())
		return err
	}

	//首次为每一个分类创建一个桶，为避免云上bucket冲突，每个bucket添加带"niwo"的前缀
	prefix := "niwo"
	if _,err := CosHandle.GetBucketStat(prefix + consts.PERSONAGE);err != nil {
		if err := CosHandle.CreateBucket(prefix + consts.PERSONAGE);err != nil {
			seelog.Errorf("create bucket %v failed, err is %v", consts.PERSONAGE, err.Error())
			return err
		}
	}

	if _,err := CosHandle.GetBucketStat(prefix + consts.SPORT);err != nil {
		if err := CosHandle.CreateBucket(prefix + consts.SPORT); err != nil {
			seelog.Errorf("create bucket %v failed, err is %v", consts.SPORT, err.Error())
			return err
		}
	}

	if _,err := CosHandle.GetBucketStat(prefix + consts.ECONOMICS);err != nil {
		if err := CosHandle.CreateBucket(prefix + consts.ECONOMICS); err != nil {
			seelog.Errorf("create bucket %v failed, err is %v", consts.SPORT, err.Error())
			return err
		}
	}

	if _,err := CosHandle.GetBucketStat(prefix + consts.MILITARY);err != nil {
		if err := CosHandle.CreateBucket(prefix + consts.MILITARY); err != nil {
			seelog.Errorf("create bucket %v failed, err is %v", consts.SPORT, err.Error())
			return err
		}
	}

	if _,err := CosHandle.GetBucketStat(prefix + consts.ENTERTAINMENT);err != nil {
		if err := CosHandle.CreateBucket(prefix + consts.ENTERTAINMENT); err != nil {
			seelog.Errorf("create bucket %v failed, err is %v", consts.SPORT, err.Error())
			return err
		}
	}

	return nil
}

func GetCosHandle() *AliOssSdk.Client{
	cosLock.RLock()
	defer cosLock.RUnlock()

	return CosHandle
}

func DestroyCosHandle() {

}