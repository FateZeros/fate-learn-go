package config

import "maple-server/pkg/logger"

func ConfigSetup(path string) {
	// 日志配置
	logger.Init()
}
