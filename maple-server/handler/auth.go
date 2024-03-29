package handler

import (
	"errors"
	"fmt"
	"maple-server/models/system"
	jwt "maple-server/pkg/jwtauth"
	"maple-server/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(system.SysUser)
		r, _ := v["role"].(system.SysRole)
		return jwt.MapClaims{
			jwt.IdentityKey: u.UserId,
			jwt.RoleIdKey:   r.RoleId,
			jwt.RoleKey:     r.RoleKey,
			jwt.NiceKey:     u.Username,
			jwt.RoleNameKey: r.RoleName,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UserName":    claims["nice"],
		"RoleKey":     claims["rolekey"],
		"UserId":      claims["identity"],
		"RoleIds":     claims["roleid"],
	}
}

// @Summary 登陆
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	var (
		loginVal system.Login
		// loginLog      system.LoginLog
		// roleValue system.SysRole
		// authUserCount int
		// addUserInfo   system.SysUser
	)

	// 获取前端过来的数据
	if err := c.ShouldBind(&loginVal); err != nil {
		// loginLog.Status = "1"
		// loginLog.Msg = "数据解析失败"
		// loginLog.Username = loginVal.Username
		// _, _ = loginLog.Create()
		return nil, jwt.ErrMissingLoginValues
	}

	user, role, e := loginVal.GetUser()
	if e == nil {
		if user.Status == "1" {
			return nil, errors.New("用户已被禁用。")
		}
		return map[string]interface{}{"user": user, "role": role}, nil
	} else {
		logger.Info(e.Error())
	}
	fmt.Printf("user: %v\n", user)
	return nil, jwt.ErrFailedAuthentication
}

func Authorizator(data interface{}, c *gin.Context) bool {

	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(system.SysUser)
		r, _ := v["role"].(system.SysRole)
		c.Set("role", r.RoleName)
		c.Set("roleIds", r.RoleId)
		c.Set("userId", u.UserId)
		c.Set("userName", u.UserName)

		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

func LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})
}
