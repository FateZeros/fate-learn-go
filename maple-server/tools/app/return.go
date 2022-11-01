package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}
