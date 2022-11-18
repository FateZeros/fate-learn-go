package system

import (
	"maple-server/global/orm"

	"github.com/pkg/errors"
)

type SysRole struct {
	RoleId   int    `json:"roleId" gorm:"primary_key;AUTO_INCREMENT"` // 角色编码
	RoleName string `json:"roleName" gorm:"type:varchar(128);"`       // 角色名称
	Status   string `json:"status" gorm:"type:int(1);"`               //
	RoleKey  string `json:"roleKey" gorm:"type:varchar(128);"`        //角色代码
	RoleSort int    `json:"roleSort" gorm:"type:int(4);"`             //角色排序
	Flag     string `json:"flag" gorm:"type:varchar(128);"`           //
	CreateBy string `json:"createBy" gorm:"type:varchar(128);"`       //
	UpdateBy string `json:"updateBy" gorm:"type:varchar(128);"`       //
	Remark   string `json:"remark" gorm:"type:varchar(255);"`         //备注
	Admin    bool   `json:"admin" gorm:"type:char(1);"`
	Params   string `json:"params" gorm:"-"`
	MenuIds  []int  `json:"menuIds" gorm:"-"`
	DeptIds  []int  `json:"deptIds" gorm:"-"`
	BaseModel
}

type MenuIdList struct {
	MenuId int `json:"menuId"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

func (role *SysRole) Get() (SysRole SysRole, err error) {
	table := orm.Eloquent.Table("sys_role")
	if role.RoleId != 0 {
		table = table.Where("role_id = ?", role.RoleId)
	}
	if role.RoleName != "" {
		table = table.Where("role_name = ?", role.RoleName)
	}
	if err = table.First(&SysRole).Error; err != nil {
		return
	}
	return
}

// 获取角色对应的菜单ids
func (role *SysRole) GetRoleMenuId() ([]int, error) {
	menuIds := make([]int, 0)
	menuList := make([]MenuIdList, 0)
	if err := orm.Eloquent.Table("sys_role_menu").Select("sys_role_menu.menu_id").Joins("LEFT JOIN sys_menu on sys_menu.menu_id=sys_role_menu.menu_id").Where("role_id = ? ", role.RoleId).Where(" sys_role_menu.menu_id not in(select sys_menu.parent_id from sys_role_menu LEFT JOIN sys_menu on sys_menu.menu_id=sys_role_menu.menu_id where role_id =? )", role.RoleId).Find(&menuList).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(menuList); i++ {
		menuIds = append(menuIds, menuList[i].MenuId)
	}
	return menuIds, nil
}

func (role *SysRole) Insert() (id int, err error) {
	i := 0
	orm.Eloquent.Table(role.TableName()).Where("(role_name = ? or role_key = ?) and `delete_time` IS NULL", role.RoleName, role.RoleKey).Count(&i)
	if i > 0 {
		return 0, errors.New("角色名称或者角色标识已经存在")
	}
	role.UpdateBy = ""
	result := orm.Eloquent.Table(role.TableName()).Create(&role)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = role.RoleId
	return
}
