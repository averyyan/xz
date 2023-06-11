package xzmongo

import (
	"sync"
	"time"

	xzcert "github.com/averyyan/xz/core/cert"
	xzcommon "github.com/averyyan/xz/core/common"
	xzservice "github.com/averyyan/xz/core/service"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOpt "go.mongodb.org/mongo-driver/mongo/options"
)

const MongoServiceType xzcommon.ServiceType = "MongoService"

type Service interface {
	xzservice.Service
	GetMongoDB(dbName string) *mongo.Database
}

type service struct {
	serviceName        xzcommon.ServiceName
	serviceDescription string
	certEntry          xzcert.Service
	client             *mongo.Client
	opts               *mongoOpt.ClientOptions
	pingTimeout        time.Duration
	insecureSkipVerify bool
	mongoDbMap         map[string]*mongo.Database
	mongoDbOpts        map[string][]*mongoOpt.DatabaseOptions
	bootstrapOnce      sync.Once
}

func (s *service) GetMongoDB(dbName string) *mongo.Database {
	return s.mongoDbMap[dbName]
}
