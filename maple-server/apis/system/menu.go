package system

import (
	"maple-server/models/system"
	"maple-server/tools"
	"maple-server/tools/app"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary 根据角色名称获取菜单列表数据（左菜单使用）
func GetMenuRole(c *gin.Context) {
	var Menu system.Menu
	result, err := Menu.SetMenuRole(tools.GetRoleName(c))

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, "")
}

// @Summary 创建菜单
func InsertMenu(c *gin.Context) {
	var data system.Menu
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	data.CreateBy = tools.GetUserIdStr(c)
	result, err := data.Create()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, "")
}

// @Summary Menu列表数据
// @Description 获取JSON
func GetMenu(c *gin.Context) {
	var data system.Menu
	id, _ := tools.StringToInt(c.Param("id"))
	data.MenuId = id
	result, err := data.GetByMenuId()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, "")
}

// @Summary 修改菜单
// @Description 获取JSON
func UpdateMenu(c *gin.Context) {
	var data system.Menu
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	data.UpdateBy = tools.GetUserIdStr(c)
	_, err = data.Update(data.MenuId)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, "", "修改成功")
}
