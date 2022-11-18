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
		v1auth.GET("/getUserInfo", system.GetUserInfo)
		v1auth.GET("/menurole", system.GetMenuRole)
		v1auth.POST("/logout", handler.LogOut)
	}
}

func RegisterMenuRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	menu := v1.Group("/menu").Use(authMiddleware.MiddlewareFunc())
	{
		menu.GET("/:id", system.GetMenu)
		menu.POST("", system.InsertMenu)
		menu.PUT("", system.UpdateMenu)
		menu.DELETE("/:id", system.DeleteMenu)
	}
}

func RegisterRoleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	role := v1.Group("/role").Use(authMiddleware.MiddlewareFunc())
	{
		role.GET("/:roleId", system.GetRole)
		role.POST("", system.InsertRole)
	}
}
