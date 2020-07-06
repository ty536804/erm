package School

import (
	erm "erm/Database"
	"time"
)

// @Summer employee info
type Employee struct {
	Id               int       `json:"id" gorm:"primary_key"`
	EmployeeName     string    `json:"employee_name" gorm:"comment:'员工姓名' " `
	EmployeeTel      string    `json:"employee_tel" gorm:"type:char(20); not null; default ''; comment:'员工电话'" `
	EmployeeAddress  string    `json:"employee_address" gorm:"type:varchar(200);not null; default ''; comment:'家庭地址' " `
	EmergencyContact string    `json:"emergency_contact" gorm:"type:varchar(100); not null; default ''; comment:'紧急联系人' " `
	EmergencyTel     string    `json:"emergency_tel" gorm:"type:char(20);not null; default ''; comment:'紧急联系人电话' " `
	EntryTime        time.Time `json:"entry_time" time_format:"2006-01-02 15:04:05" gorm:"comment:'入职时间' " `
	IsTeach          int       `json:"is_teach" gorm:"default '0'; comment:'是否授课 0否 1是' " `
	ERemark          string    `json:"e_remark" gorm:"type:text; comment:'备注'" `
	RuleId           int       `json:"rule_id" gorm:"comment:'权限' " `
	Sex              int       `json:"sex" gorm:"comment:'性别 0男 1女' " `

	erm.Model
}
