package xzmongo

import (
	"time"

	xzcommon "github.com/averyyan/xz/core/common"
	xzservice "github.com/averyyan/xz/core/service"
	xzutil "github.com/averyyan/xz/core/util"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOpt "go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterServiceYAML(raw []byte) map[xzcommon.ServiceName]xzservice.Service {
	var config Config
	entryMap := make(map[xzcommon.ServiceName]xzservice.Service)
	if err := xzutil.UnmarshalConfigYAML(raw, &config); err != nil {
		panic(err.Error())
	}
	for i := range config.Mongo {
		element := config.Mongo[i]
		if !element.Enabled {
			continue
		}
		clientOpt := ToClientOptions(element)
		opts := []Option{
			WithName(xzcommon.ServiceName(element.Name)),
			WithDescription(element.Description),
			WithClientOptions(clientOpt),
			WithCertEntry(element.CertEntry),
			WithPingTimeout(element.PingTimeout),
			WithInsecureSkipVerify(element.InsecureSkipVerify),
		}
		// iterate database
		for i := range element.Database {
			opts = append(opts, WithDatabase(element.Database[i].Name))
		}
		serv := registerService(opts...)
		entryMap[serv.serviceName] = serv
	}
	return entryMap
}

func registerService(opts ...Option) *service {
	serv := &service{
		serviceName:        "Mongo服务",
		serviceDescription: "Mongo服务",
		mongoDbMap:         make(map[string]*mongo.Database),
		mongoDbOpts:        make(map[string][]*mongoOpt.DatabaseOptions),
		pingTimeout:        3 * time.Second,
	}
	clientOptions := &mongoOpt.ClientOptions{
		AppName: (*string)(&serv.serviceName),
	}
	serv.opts = clientOptions
	for i := range opts {
		opts[i](serv)
	}
	return serv
}

func ToClientOptions(config *ConfigE) *mongoOpt.ClientOptions {
	if config == nil {
		return &mongoOpt.ClientOptions{}
	}
	opt := &mongoOpt.ClientOptions{
		Hosts:      config.Hosts,
		ReplicaSet: &config.ReplicaSet,
	}
	// auth
	if config.Auth != nil {
		opt.Auth = &mongoOpt.Credential{
			Username: config.Auth.Username,
			Password: config.Auth.Password,
		}
	}
	return opt
}
