package router

import (
	systemRouter "maple-server/router/system"

	"github.com/gin-gonic/gin"
)

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")

	systemRouter.SysBaseRouter(g)

	// 无需认证
	systemRouter.SysNoCheckRoleRouter(g)

	return g
}
