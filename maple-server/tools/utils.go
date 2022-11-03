package tools

import (
	"maple-server/pkg/logger"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func HasError(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		if msg == "" {
			msg = err.Error()
		}
		logger.Info(err)
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		logger.Info(err.Error())
		return false, err
	}
	return true, nil
}
