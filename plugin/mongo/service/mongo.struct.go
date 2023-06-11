package xzmongoservice

import (
	"fmt"

	xzcommon "github.com/averyyan/xz/core/common"
	xzmongo "github.com/averyyan/xz/plugin/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBService struct {
	ServiceName string `yaml:"serviceName"`
	Database    string `yaml:"database"`
	Collection  string `yaml:"collection"`
}

// 获取数据库连接
func (mongoDBService *MongoDBService) GetCollection() *mongo.Collection {
	mongodb := xzmongo.GetMongoDB(
		xzcommon.ServiceName(mongoDBService.ServiceName),
		mongoDBService.Database,
	)
	if mongodb == nil {
		return nil
	}
	return mongodb.Collection(mongoDBService.Collection)
}

// 检查
func (mongoDBService *MongoDBService) CheckCollection() error {
	if mongoDBService.GetCollection() == nil {
		return fmt.Errorf(
			"数据库服务【%s-%s-%s】未存在",
			mongoDBService.ServiceName,
			mongoDBService.Database,
			mongoDBService.Collection,
		)
	}
	return nil
}
