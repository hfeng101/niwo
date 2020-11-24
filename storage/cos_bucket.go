package storage

import (
	"github.com/cihub/seelog"
	"io"
)

func StorageFromIoReader (catalog string, objectKey string, file io.Reader)error{
	bucketHandle,err := CosHandle.Bucket(catalog)
	if err != nil {
		seelog.Errorf("get bucketHandle %v failed, err is s%v", catalog,err.Error())
		return err
	}

	if err := bucketHandle.PutObject(objectKey, file); err != nil {
		seelog.Errorf("PutObject failed, err is %v", err.Error())
		return err
	}

	return nil
}

func GetObjectFromKey (catalog string, objectKey string) (io.Reader, error) {
	bucketHandle,err := CosHandle.Bucket(catalog)
	if err != nil {
		seelog.Errorf("get bucketHandle %v failed, err is s%v", catalog,err.Error())
		return nil, err
	}

	data,err := bucketHandle.GetObject(objectKey)
	if err != nil {
		seelog.Errorf("Get object failed, err is %v", err.Error())
		return nil,nil
	}

	return data, nil
}