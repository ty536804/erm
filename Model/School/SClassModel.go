package School

import (
	erm "erm/Database"
)

// @Summer class name
type ClassName struct {
	Id       int    `json:"id" gorm:"primary_key"`
	SName    string `json:"s_name" gorm:"type:varchar(100) ;not null; default ''; comment:'班级名称'"`
	CourseId int    `json:"course_id" gorm:"comment:'课程' " `
	Schedule int    `json:"schedule" gorm:"comment:'是否排课' " `
	SStatus  int    `json:"s_status" gorm:"comment:'结业状态' " `
	SRemark  string `json:"s_remark" gorm:"comment:'备注' " `
	SchoolId int    `json:"school_id" gorm:"index:school_id" `

	erm.Model
}
