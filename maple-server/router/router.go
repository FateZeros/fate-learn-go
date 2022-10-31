package router

import (
	systemRouter "maple-server/router/system"

	"github.com/gin-gonic/gin"
)

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")

	systemRouter.SysBaseRouter(g)

	return g
}
