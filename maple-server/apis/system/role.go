package system

import (
	"maple-server/models/system"
	"maple-server/tools"
	"maple-server/tools/app"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary 获取 role 数据
func GetRole(c *gin.Context) {
	var (
		err  error
		Role system.SysRole
	)
	Role.RoleId, _ = tools.StringToInt(c.Param("roleId"))

	result, err := Role.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	menuIds, err := Role.GetRoleMenuId()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	result.MenuIds = menuIds
	app.OK(c, result, "")
}

// @Summary 创建角色
func InsertRole(c *gin.Context) {
	var data system.SysRole
	data.CreateBy = tools.GetUserIdStr(c)
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	id, err := data.Insert()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	data.RoleId = id
	var menu system.RoleMenu
	_, err = menu.Insert(id, data.MenuIds)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, data, "添加角色成功")
}
