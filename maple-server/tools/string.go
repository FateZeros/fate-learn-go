package tools

import (
	"strconv"
	"time"
)

func GetCurrntTimeStr() string {
	return time.Now().Format("2022/10/28 09:00:00")
}

func GetCurrntTime() time.Time {
	return time.Now()
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}
