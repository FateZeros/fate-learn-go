package config

import (
	"io/ioutil"
	"maple-server/pkg/logger"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cfgApplication *viper.Viper
var cfgDatabase *viper.Viper

// 载入配置文件
func ConfigSetup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)

	if err != nil {
		// logger.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	// Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		// logger.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	// 数据库初始化
	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("config not found settings.database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	// 启动参数
	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("config not found settings.application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	// 日志配置
	logger.Init()

}

func SetConfig(configPath string, key string, value interface{}) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	_ = viper.WriteConfig()
}
