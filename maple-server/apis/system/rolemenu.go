package system

import (
	"maple-server/tools/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertRoleMenu(c *gin.Context) {
	var res app.Response
	res.Msg = "添加成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}
