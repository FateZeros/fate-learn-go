package tools

import (
	"maple-server/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

func GetUserId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return int((data["identity"]).(float64))
	}
	logger.Info("********** 路径：" + c.Request.URL.Path + "  请求方法：" + c.Request.Method + "  说明：缺少identity")
	return 0
}
