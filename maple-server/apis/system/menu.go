package system

import (
	"maple-server/models/system"
	"maple-server/tools"
	"maple-server/tools/app"

	"github.com/gin-gonic/gin"
)

func GetMenuRole(c *gin.Context) {
	var Menu system.Menu
	result, err := Menu.SetMenuRole(tools.GetRoleName(c))

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, "")

}
