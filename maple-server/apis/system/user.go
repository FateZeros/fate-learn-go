package system

import (
	"maple-server/models/system"
	"maple-server/tools"
	"maple-server/tools/app"

	"github.com/gin-gonic/gin"
)

// @Summary 摘要 比如获取用户列表
// @Schemes
// @Description 这里写描述 get users
func GetUserInfo(c *gin.Context) {
	var mp = make(map[string]interface{})

	sysuser := system.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	user, err := sysuser.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	mp["userName"] = user.NickName
	mp["userId"] = user.UserId
	mp["deptId"] = user.DeptId
	mp["name"] = user.NickName

	app.OK(c, mp, "")
}
