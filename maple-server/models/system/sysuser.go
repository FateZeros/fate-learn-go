package system

import "maple-server/global/orm"

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
}

type SysUserB struct {
	NickName string `gorm:"type:varchar(128)" json:"nickName"` // 昵称
	Phone    string `gorm:"type:varchar(11)" json:"phone"`     // 手机号
	RoleId   int    `gorm:"type:int(11)" json:"roleId"`        // 角色编码
	Salt     string `gorm:"type:varchar(255)" json:"salt"`     //盐
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`   //头像
	Sex      string `gorm:"type:varchar(255)" json:"sex"`      //性别
	Email    string `gorm:"type:varchar(128)" json:"email"`    //邮箱
	DeptId   int    `gorm:"type:int(11)" json:"deptId"`        //部门编码
	PostId   int    `gorm:"type:int(11)" json:"postId"`        //职位编码
	CreateBy string `gorm:"type:varchar(128)" json:"createBy"` //
	UpdateBy string `gorm:"type:varchar(128)" json:"updateBy"` //
	Remark   string `gorm:"type:varchar(255)" json:"remark"`   //备注
	Status   string `gorm:"type:int(1);" json:"status"`
	Params   string `gorm:"-" json:"params"`
	BaseModel
}

type LoginM struct {
	UserName
	PassWord
}

type UserName struct {
	Username string `gorm:"type:varchar(64)" json:"username"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"type:varchar(128)" json:"password"`
}

type SysUser struct {
	SysUserId
	SysUserB
	LoginM
}

type SysUserView struct {
	SysUserId
	SysUserB
	LoginM
	RoleName string `gorm:"column:role_name"  json:"role_name"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

// 获取用户数据
func (e *SysUser) Get() (SysUserView SysUserView, err error) {

	table := orm.Eloquent.Table(e.TableName()).Select([]string{"sys_user.*", "sys_role.role_name"})
	table = table.Joins("left join sys_role on sys_user.role_id=sys_role.role_id")
	if e.UserId != 0 {
		table = table.Where("user_id = ?", e.UserId)
	}

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	if e.RoleId != 0 {
		table = table.Where("role_id = ?", e.RoleId)
	}

	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}

	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if err = table.First(&SysUserView).Error; err != nil {
		return
	}

	SysUserView.Password = ""
	return
}
