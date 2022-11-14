package tools

import (
	"fmt"
	jwt "maple-server/pkg/jwtauth"
	"maple-server/pkg/logger"

	"github.com/gin-gonic/gin"
)

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

func GetUserIdStr(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return Int64ToString(int64((data["identity"]).(float64)))
	}
	logger.Info("********** 路径：" + c.Request.URL.Path + "  请求方法：" + c.Request.Method + "  缺少identity")
	return ""
}

func GetUserId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return int((data["identity"]).(float64))
	}
	logger.Info("********** 路径：" + c.Request.URL.Path + "  请求方法：" + c.Request.Method + "  说明：缺少identity")
	return 0
}

func GetRoleName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["rolekey"] != nil {
		return (data["rolekey"]).(string)
	}
	fmt.Println("********** 路径：" + c.Request.URL.Path + "  请求方法：" + c.Request.Method + "  缺少rolekey")
	return ""
}
