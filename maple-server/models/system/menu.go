package system

import (
	"errors"
	"maple-server/global/orm"
	"maple-server/tools"
)

type Menu struct {
	MenuId     int    `json:"menuId" gorm:"primary_key;AUTO_INCREMENT"`
	MenuName   string `json:"menuName" gorm:"type:varchar(128);"`
	Title      string `json:"title" gorm:"type:varchar(64);"`
	Icon       string `json:"icon" gorm:"type:varchar(128);"`
	Path       string `json:"path" gorm:"type:varchar(128);"`
	Paths      string `json:"paths" gorm:"type:varchar(128);"`
	MenuType   string `json:"menuType" gorm:"type:varchar(1);"`
	Action     string `json:"action" gorm:"type:varchar(16);"`
	Permission string `json:"permission" gorm:"type:varchar(32);"`
	ParentId   int    `json:"parentId" gorm:"type:int(11);"`
	NoCache    bool   `json:"noCache" gorm:"type:char(1);"`
	Breadcrumb string `json:"breadcrumb" gorm:"type:varchar(255);"`
	Component  string `json:"component" gorm:"type:varchar(255);"`
	Sort       int    `json:"sort" gorm:"type:int(4);"`
	Visible    string `json:"visible" gorm:"type:char(1);"`
	CreateBy   string `json:"createBy" gorm:"type:varchar(128);"`
	UpdateBy   string `json:"updateBy" gorm:"type:varchar(128);"`
	IsFrame    string `json:"isFrame" gorm:"type:int(1);DEFAULT:0;"`
	Params     string `json:"params" gorm:"-"`
	RoleId     int    `gorm:"-"`
	Children   []Menu `json:"children" gorm:"-"`
	IsSelect   bool   `json:"is_select" gorm:"-"`
	BaseModel
}

type Menus struct {
	MenuId     int    `json:"menuId" gorm:"column:menu_id;primary_key"`
	MenuName   string `json:"menuName" gorm:"column:menu_name"`
	Title      string `json:"title" gorm:"column:title"`
	Icon       string `json:"icon" gorm:"column:icon"`
	Path       string `json:"path" gorm:"column:path"`
	MenuType   string `json:"menuType" gorm:"column:menu_type"`
	Action     string `json:"action" gorm:"column:action"`
	Permission string `json:"permission" gorm:"column:permission"`
	ParentId   int    `json:"parentId" gorm:"column:parent_id"`
	NoCache    bool   `json:"noCache" gorm:"column:no_cache"`
	Breadcrumb string `json:"breadcrumb" gorm:"column:breadcrumb"`
	Component  string `json:"component" gorm:"column:component"`
	Sort       int    `json:"sort" gorm:"column:sort"`

	Visible  string `json:"visible" gorm:"column:visible"`
	Children []Menu `json:"children" gorm:"-"`

	CreateBy string `json:"createBy" gorm:"column:create_by"`
	UpdateBy string `json:"updateBy" gorm:"column:update_by"`
	Params   string `json:"params" gorm:"-"`
	BaseModel
}

type MenuLabel struct {
	Id       int         `json:"id" gorm:"-"`
	Label    string      `json:"label" gorm:"-"`
	Children []MenuLabel `json:"children" gorm:"-"`
}

type MenuRole struct {
	Menus
	IsSelect bool `json:"is_select" gorm:"-"`
}

func (Menu) TableName() string {
	return "sys_menu"
}

func FindMenuLabel(menulist *[]Menu, menu MenuLabel) MenuLabel {
	list := *menulist

	min := make([]MenuLabel, 0)
	for j := 0; j < len(list); j++ {
		if menu.Id != list[j].ParentId {
			continue
		}
		mi := MenuLabel{}
		mi.Id = list[j].MenuId
		mi.Label = list[j].Title
		mi.Children = []MenuLabel{}
		if list[j].MenuType != "F" {
			ms := FindMenuLabel(menulist, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

func FindMenu(menulist *[]Menu, menu Menu) Menu {
	list := *menulist

	min := make([]Menu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := Menu{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Action = list[j].Action
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.NoCache = list[j].NoCache
		mi.Breadcrumb = list[j].Breadcrumb
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Visible = list[j].Visible
		mi.CreatedAt = list[j].CreatedAt
		mi.UpdatedAt = list[j].UpdatedAt
		mi.Children = []Menu{}

		if mi.MenuType != "F" {
			ms := FindMenu(menulist, mi)
			min = append(min, ms)

		} else {
			min = append(min, mi)
		}

	}
	menu.Children = min
	return menu
}

func (e *Menu) SetMenuRole(rolename string) (m []Menu, err error) {
	menulist, err := e.GetByRoleName(rolename)

	m = make([]Menu, 0)
	for i := 0; i < len(menulist); i++ {
		if menulist[i].ParentId != 0 {
			continue
		}
		menusInfo := FindMenu(&menulist, menulist[i])

		m = append(m, menusInfo)
	}
	return
}

func (e *Menu) GetByRoleName(rolename string) (Menus []Menu, err error) {
	table := orm.Eloquent.Table(e.TableName()).Select("sys_menu.*").Joins("left join sys_role_menu on sys_role_menu.menu_id=sys_menu.menu_id")
	table = table.Where("sys_role_menu.role_name=? and menu_type in ('M','C')", rolename)
	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}

func (e *Menu) Create() (id int, err error) {
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	err = InitPath(e)
	if err != nil {
		return
	}
	id = e.MenuId
	return
}

func InitPath(menu *Menu) (err error) {
	parentMenu := new(Menu)
	if int(menu.ParentId) != 0 {
		orm.Eloquent.Table("sys_menu").Where("menu_id = ?", menu.ParentId).First(parentMenu)
		if parentMenu.Paths == "" {
			err = errors.New("父级paths异常，请尝试对当前节点父级菜单进行更新操作！")
			return
		}
		menu.Paths = parentMenu.Paths + "/" + tools.IntToString(menu.MenuId)
	} else {
		menu.Paths = "/0/" + tools.IntToString(menu.MenuId)
	}
	orm.Eloquent.Table("sys_menu").Where("menu_id = ?", menu.MenuId).Update("paths", menu.Paths)
	return
}

func (e *Menu) GetByMenuId() (Menu Menu, err error) {
	table := orm.Eloquent.Table(e.TableName())
	table = table.Where("menu_id = ?", e.MenuId)
	if err = table.Find(&Menu).Error; err != nil {
		return
	}
	return
}

func (e *Menu) Update(id int) (update Menu, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Update(&e).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Update(&e).Error; err != nil {
		return
	}
	err = InitPath(e)
	if err != nil {
		return
	}
	return
}

func (e *Menu) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("menu_id = ?", id).Delete(&Menu{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
