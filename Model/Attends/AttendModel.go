package Attends

import (
	erm "erm/Database"
	"time"
)

// todo
type AttendLog struct {
	Id           int       `json:"id" gorm:"primary_key:true;unique"`
	AttendTime   time.Time `json:"attend_time" gorm:"comment:'点名时间'"`
	SName        string    `json:"s_name" gorm:"type:varchar(100) ;not null; default ''; comment:'班级名称'"`
	CourseName   string    `json:"course_name" gorm:"type:varchar(100); not null default: '';comment:'课程名称'" `
	EmployeeName string    `json:"employee_name" gorm:"comment:'上课老师' " `
	erm.Model
}
