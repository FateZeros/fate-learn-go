package system

import (
	"maple-server/apis/system"
	"maple-server/handler"
	jwt "maple-server/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func SysBaseRouter(r *gin.RouterGroup) {
	r.GET("/hello", handler.Ping)
}

func SysNoCheckRoleRouter(r *gin.RouterGroup) {

}

func RegisterBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		v1auth.GET("/getinfo", system.GetUserInfo)
	}
}
