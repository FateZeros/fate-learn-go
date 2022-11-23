package system

import (
	"maple-server/global/orm"
	"maple-server/tools"
)

type Dept struct {
	DeptId   int    `json:"deptId" gorm:"primary_key;AUTO_INCREMENT"` //部门编码
	ParentId int    `json:"parentId" gorm:"type:int(11);"`            //上级部门
	DeptPath string `json:"deptPath" gorm:"type:varchar(255);"`       //
	DeptName string `json:"deptName"  gorm:"type:varchar(128);"`      //部门名称
	Sort     int    `json:"sort" gorm:"type:int(4);"`                 //排序
	Leader   int    `json:"leader" gorm:"type:int(11);"`              //负责人
	Phone    string `json:"phone" gorm:"type:varchar(11);"`           //手机
	Email    string `json:"email" gorm:"type:varchar(64);"`           //邮箱
	Status   string `json:"status" gorm:"type:int(1);"`               //状态
	CreateBy string `json:"createBy" gorm:"type:varchar(64);"`
	UpdateBy string `json:"updateBy" gorm:"type:varchar(64);"`
	Params   string `json:"params" gorm:"-"`
	Children []Dept `json:"children" gorm:"-"`
	BaseModel
}

func (Dept) TableName() string {
	return "sys_dept"
}

type DeptLabel struct {
	Id       int         `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []DeptLabel `gorm:"-" json:"children"`
}

func (dept *Dept) Create() (Dept, error) {
	var doc Dept
	result := orm.Eloquent.Table(dept.TableName()).Create(&dept)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	deptPath := "/" + tools.IntToString(dept.DeptId)
	if int(dept.DeptId) != 0 {
		var deptParent Dept
		orm.Eloquent.Table(dept.TableName()).Where("dept_id = ?", dept.ParentId).First(&deptParent)
		deptPath = deptParent.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	var mp = map[string]string{}
	mp["deptPath"] = deptPath
	if err := orm.Eloquent.Table(dept.TableName()).Where("dept_id = ?", dept.DeptId).Update(mp).Error; err != nil {
		err := result.Error
		return doc, err
	}
	doc = *dept
	doc.DeptPath = deptPath
	return doc, nil
}

func (dept *Dept) Get() (Dept, error) {
	var doc Dept
	table := orm.Eloquent.Table(dept.TableName())
	if dept.DeptId != 0 {
		table = table.Where("dept_id = ?", dept.DeptId)
	}
	if dept.DeptName != "" {
		table = table.Where("dept_name = ?", dept.DeptName)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
