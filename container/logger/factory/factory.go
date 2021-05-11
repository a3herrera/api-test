package factory

import "github.com/a3herrera/api-test/container/logger/factory/zap"

var logFactoryBuilderMap = map[string]logFactoryInterface{
	"zap": &zap.ZapFactory{},
}

type logFactoryInterface interface {
	Build() error
}

func GetLogFactoryBuilder(key string) logFactoryInterface {
	return logFactoryBuilderMap[key]
}
