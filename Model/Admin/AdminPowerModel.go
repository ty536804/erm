package Admin

import (
	erm "erm/Database"
	"erm/Pkg/Setting"
	"fmt"
)

// @Summer 权限表
type SysAdminPower struct {
	Id        int    `json:"id" gorm:"primary_key:true;unique;"`
	PowerName string `json:"power_name" gorm:"comment:'权限名称' " `
	Level     int    `json:"level" gorm:"not null;default 0;comment:'级别'"`
	Pid       int    `json:"pid" gorm:"not null;default 0;comment:'父ID' " `
	Status    int    `json:"status" gorm:"comment:'状态 1有效 0无效'" `
	Desc      string `json:"desc" gorm:"type:varchar(200); not null; default '';comment:'描述'" `

	erm.Model
}

// @Summer 添加权限
func AddPower(data map[string]interface{}) (isOk bool) {
	powerRes := erm.Db.Create(&SysAdminPower{
		PowerName: data["power_name"].(string),
		Pid:       data["pid"].(int),
		Level:     data["pid"].(int),
		Status:    data["status"].(int),
		Desc:      data["desc"].(string),
	})
	if powerRes.Value != nil {
		fmt.Printf("添加权限失败:%s", powerRes.Value)
		return false
	}
	return
}

// @Summer 编辑权限
func EditPower(id int, data map[string]interface{}) (power SysAdminPower) {
	erm.Db.Where("id = ?", id).Update(&power)
	return
}

// @Summer 获取权限列表
func GetPower(page int) (power []SysAdminPower) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * Setting.PageSize
	}
	erm.Db.Limit(Setting.PageSize).Offset(offset).Find(&power)
	return
}
