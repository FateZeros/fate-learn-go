package app

import (
	"maple-server/pkg/logger"
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

func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	res.Msg = err.Error()
	if msg != "" {
		res.Msg = msg
	}
	logger.Error(res.Msg)
	c.JSON(http.StatusOK, res.ReturnError(code))
}
