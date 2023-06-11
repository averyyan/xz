package xzmongo

import (
	xzcommon "github.com/averyyan/xz/core/common"
	xzcontext "github.com/averyyan/xz/core/context"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetService(name xzcommon.ServiceName) *service {
	if v := xzcontext.GetService(MongoServiceType, name); v != nil {
		if res, ok := v.(*service); ok {
			return res
		}
	}
	return nil
}

func GetMongoDB(name xzcommon.ServiceName, dbName string) *mongo.Database {
	if serv := GetService(name); serv != nil {
		return serv.GetMongoDB(dbName)
	}

	return nil
}
