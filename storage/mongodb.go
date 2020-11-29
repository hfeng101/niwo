package storage

import (
	"context"
	"github.com/cihub/seelog"
	"github.com/hfeng101/niwo/utils/config"
	"github.com/hfeng101/niwo/utils/consts"

	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDbClient *mongo.Client
	MongoDbHandle *mongo.Database
	MongoDbCollectionHandle *mongo.Collection

	MongoLock = sync.RWMutex{}
)

func InitMongoDb() error{
	ctx , cancel :=context.WithTimeout(context.Background(),10*time.Second)
	//养成良好的习惯，在调用WithTimeout之后defer cancel()
	defer cancel()

	var err error
	MongoDbClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().Other["mongodb-dev"].(string)));
	if  err != nil {
		seelog.Errorf("Connect to mongodb failed, err is %v",err.Error())
		return err
	}

	//务必保证数据库已经创建
	MongoDbHandle = MongoDbClient.Database(config.GetConfig().Other["mongo-dev-database"].(string))
	//MongoDbCollectionHandle = MongoDbHandle.Database(config.GetConfig().Other["mongo-dev-storage"].(string)).Collection(config.GetConfig().Other["mongo-dev-collection"].(string))
	//MongoDbCollectionHandle = MongoDbClient.Database(config.GetConfig().Other["mongo-dev-storage"].(string)).Collection(config.GetConfig().Other["mongo-dev-collection"].(string))
	collections := []string{
		consts.PERSONAGE,
		consts.SPORT,
		consts.ECONOMICS,
		consts.MILITARY,
		consts.ENTERTAINMENT,
	}

	if err := CreateMongoCollectionAndIndex(collections);err != nil {
		seelog.Errorf("CreateMongoIndexes failed, err is %v",err.Error())
		return err
	}

	return nil
}

func GetMongoDbClient() *mongo.Client{
	MongoLock.RLock()
	defer MongoLock.RUnlock()
	return MongoDbClient
}

func GetMongoDbHandle() *mongo.Database{
	MongoLock.RLock()
	defer MongoLock.RUnlock()
	return MongoDbHandle
}

func GetMongoDbCollectionHandle() *mongo.Collection{
	MongoLock.RLock()
	defer MongoLock.RUnlock()
	return MongoDbCollectionHandle
}

func CreateMongoCollectionAndIndex(collections []string) error {
	seelog.Infof("mongo.Connect successed, MongoDbHandle is %v", *MongoDbHandle)
	for _,collection := range(collections) {
		//if MongoDbHandle.Collection(collection).

		//创建collection
		niwoCollection := MongoDbHandle.Collection(collection)
		seelog.Infof("niwoCollection is %v", *niwoCollection)
		//MongoDbHandle.CreateCollection()
		expireIn := int32(360000000)
		_, err := niwoCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
			{
				Keys:    bson.M{"uuid": 1},
				Options: options.Index().SetBackground(true).SetExpireAfterSeconds(expireIn),
			},
		})
		if err != nil {
			seelog.Errorf("CreateMongoCollection %v failed, err is %v",collection, err.Error())
			return err
		}
	}

	return nil
}
