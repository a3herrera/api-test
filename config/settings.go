package config

import (
	conf "github.com/spf13/viper"
)

func init() {
	setSettings()
}

func setSettings() {
	defaultAppConfig := map[string]interface{}{
		AppName: "api test",

		LoggerCode:         "zap",
		LoggerLevel:        "debug",
		LoggerEnableCaller: true,

		LoggerZapEncoding:                "console",
		LoggerZapDevelopment:             true,
		LoggerZapEncoderConfigKeyMessage: "msg",
		LoggerZapEncoderConfigKeyLevel:   "level",
		LoggerZapEncoderConfigKeyTime:    "ts",
		LoggerZapEncoderConfigKeyName:    "logger",
		LoggerZapEncoderConfigKeyCaller:  "",
	}

	for key, value := range defaultAppConfig {
		conf.SetDefault(key, value)
	}
}
