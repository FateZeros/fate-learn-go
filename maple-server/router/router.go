package router

import (
	"maple-server/apis/tpl"
	"maple-server/pkg/jwtauth"
	systemRouter "maple-server/router/system"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	systemRouter.SysBaseRouter(g)

	// 静态文件
	sysStaticFileRouter(g, r)

	// swagger；注意：生产环境可以注释掉
	sysSwaggerRouter(g)

	// 无需认证
	// systemRouter.SysNoCheckRoleRouter(g)

	sysCheckRoleRouterInit(g, authMiddleware)

	return g
}

func sysStaticFileRouter(r *gin.RouterGroup, g *gin.Engine) {
	r.Static("/static", "./static")
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	r.POST("/login", authMiddleware.LoginHandler)

	v1 := r.Group("/api/v1")

	// 兼容前后端不分离的情况
	r.GET("/", tpl.Tpl)

	// 系统管理
	systemRouter.RegisterBaseRouter(v1, authMiddleware)
	// 菜单
	systemRouter.RegisterMenuRouter(v1, authMiddleware)
	// 角色
	systemRouter.RegisterRoleRouter(v1, authMiddleware)
	// 部门
	systemRouter.RegisterDeptRouter(v1, authMiddleware)
}
