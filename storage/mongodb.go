package storage

import (
	"context"
	"github.com/cihub/seelog"
	"github.com/hfeng101/niwo/utils/config"
	"github.com/hfeng101/niwo/utils/consts"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDbClient *mongo.Client
	MongoDbHandle *mongo.Database
	MongoDbCollectionHandle *mongo.Collection

	MongoLock = sync.RWMutex{}
)

func InitMongoDb() {
	ctx , cancel :=context.WithTimeout(context.Background(),10*time.Second)
	//养成良好的习惯，在调用WithTimeout之后defer cancel()
	defer cancel()

	MongoDbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().Other["mongodb-dev"].(string)));
	if  err != nil {
		MongoDbClient = nil
		MongoDbHandle = nil
		MongoDbCollectionHandle = nil
		seelog.Errorf("Connect to mongodb failed, err is %v",err.Error())
		return
	}

	//待定
	MongoDbHandle = MongoDbClient.Database(config.GetConfig().Other["mongo-dev-storage"].(string))
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
		return
	}
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
	for _,collection := range(collections) {
		//创建collection
		niwoCollection := MongoDbHandle.Collection(collection)

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
