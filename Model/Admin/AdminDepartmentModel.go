package Admin

import (
	erm "erm/Database"
)

// @Summer 部门表
type SysAdminDepartment struct {
	Id        int    `json:"id" gorm:"primary_key:true;unique"`
	PowerName string `json:"power_name" gorm:"comment:'部门名称' " `
	DPid      int    `json:"d_pid" gorm:"default 0; comment:'父ID' " `
	DStatus   int    `json:"d_status" gorm:"default 0; comment:'状态 1有效 0无效'" `
	Desc      string `json:"desc" gorm:"type:varchar(200); not null; default '';comment:'描述'" `

	erm.Model
}
