package router

import (
	"maple-server/handler"
	"maple-server/middleware"
	"maple-server/tools"
	config2 "maple-server/tools/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	if config2.ApplicationConfig.IsHttps {
		r.Use(handler.TlsHandler())
	}
	middleware.InitMiddleware(r)

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	tools.HasError(err, "JWT Init Error", 500)

	// 注册系统路由
	InitSysRouter(r, authMiddleware)

	return r
}
