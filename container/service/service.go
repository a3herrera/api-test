package service

import (
	"github.com/a3herrera/api-test/config"
	"github.com/a3herrera/api-test/container/logger/factory"
	"github.com/pkg/errors"
	conf "github.com/spf13/viper"
)

func InitApp() error {
	if err := loadLogger(); err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func loadLogger() error {
	loggerCode := conf.GetString(config.LoggerCode)
	if err := factory.GetLogFactoryBuilder(loggerCode).Build(); err != nil {
		return errors.Wrap(err, "fail to load logger")
	}

	return nil
}
