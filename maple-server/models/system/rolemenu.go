package system

import "maple-server/global/orm"

type RoleMenu struct {
	RoleId   int    `gorm:"type:int(11)"`
	MenuId   int    `gorm:"type:int(11)"`
	RoleName string `gorm:"type:varchar(128)"`
	CreateBy string `gorm:"type:varchar(128)"`
	UpdateBy string `gorm:"type:varchar(128)"`
	Id       int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
}

func (RoleMenu) TableName() string {
	return "sys_role_menu"
}

func (roleMenu *RoleMenu) Insert(roleId int, menuId []int) (bool, error) {
	var role SysRole
	if err := orm.Eloquent.Table("sys_role").Where("role_id = ?", roleId).First(&role).Error; err != nil {
		return false, err
	}
	var menu []Menu
	if err := orm.Eloquent.Table("sys_menu").Where("menu_id in (?)", menuId).Find(&menu).Error; err != nil {
		return false, err
	}
	return true, nil
}
