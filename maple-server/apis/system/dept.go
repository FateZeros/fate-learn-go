package system

import (
	"fmt"
	"maple-server/models/system"
	"maple-server/tools"
	"maple-server/tools/app"
	"maple-server/tools/app/msg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary 部门列表数据
// @Description 获取JSON
func GetDept(c *gin.Context) {
	var (
		err  error
		Dept system.Dept
	)
	fmt.Printf("c.Param(\"deptId\"): %v\n", c.Param("deptId"))
	Dept.DeptId, _ = tools.StringToInt(c.Param("deptId"))

	result, err := Dept.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, msg.GetSuccess)
}

func InsertDept(c *gin.Context) {
	var dept system.Dept
	err := c.BindWith(&dept, binding.JSON)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	dept.CreateBy = tools.GetUserIdStr(c)
	result, err := dept.Create()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, msg.CreatedSuccess)
}
