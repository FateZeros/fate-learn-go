package router

import (
	"maple-server/handler"
	"maple-server/middleware"
	config2 "maple-server/tools/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	if config2.ApplicationConfig.IsHttps {
		r.Use(handler.TlsHandler())
	}
	middleware.InitMiddleware(r)

	// 注册系统路由
	InitSysRouter(r)

	return r
}
