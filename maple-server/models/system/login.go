package system

import (
	"fmt"
	"maple-server/global/orm"
	"maple-server/tools"
)

type Login struct {
	Username  string `form:"UserName" json:"username" binding:"required"`
	Password  string `form:"Password" json:"password" binding:"required"`
	Code      string `form:"Code" json:"code"`
	UUID      string `form:"UUID" json:"uuid"`
	LoginType int    `form:"LoginType" json:"loginType"`
}

func (u *Login) GetUser() (user SysUser, e error) {
	fmt.Printf("u.Username: %v\n", u.Username)
	e = orm.Eloquent.Table("sys_user").Where("username = ? ", "admin").Find(&user).Error
	fmt.Printf("e: %v\n", e)
	if e != nil {
		return
	}

	fmt.Printf("u.LoginType: %v\n", u.LoginType)
	// 验证密码
	if u.LoginType == 0 {
		_, e = tools.CompareHashAndPassword(user.Password, u.Password)
		if e != nil {
			return
		}
	}

	return
}
