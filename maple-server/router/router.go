package router

import (
	"maple-server/pkg/jwtauth"
	systemRouter "maple-server/router/system"

	"github.com/gin-gonic/gin"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	systemRouter.SysBaseRouter(g)

	// 无需认证
	// systemRouter.SysNoCheckRoleRouter(g)

	sysCheckRoleRouterInit(g, authMiddleware)

	return g
}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	r.POST("/login", authMiddleware.LoginHandler)

	v1 := r.Group("/api/v1")

	// 系统管理
	systemRouter.RegisterBaseRouter(v1, authMiddleware)
}
