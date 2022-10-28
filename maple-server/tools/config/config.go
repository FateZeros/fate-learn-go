package config

import (
	"maple-server/pkg/logger"

	"github.com/spf13/viper"
)

// 载入配置文件
func ConfigSetup(path string) {
	// 日志配置
	logger.Init()
}

func SetConfig(configPath string, key string, value interface{}) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	_ = viper.WriteConfig()
}
