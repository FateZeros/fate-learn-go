package tools

import "time"

func GetCurrntTimeStr() string {
	return time.Now().Format("2022/10/28 09:00:00")
}
