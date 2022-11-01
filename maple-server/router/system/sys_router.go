package system

import (
	"maple-server/handler"

	"github.com/gin-gonic/gin"
)

func SysBaseRouter(r *gin.RouterGroup) {
	r.GET("/hello", handler.Ping)
}

func SysNoCheckRoleRouter(r *gin.RouterGroup) {
	// v1 := r.Group("/api/v1")

}
