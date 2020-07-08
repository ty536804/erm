package Courses

import (
	erm "erm/Database"
	"time"
)

// @Summer  Course info
type Course struct {
	Id             int       `json:"id" gorm:"primary_key:true;unique"`
	CourseName     string    `json:"course_name" gorm:"type:varchar(100); not null default: '';comment:'课程名称'" `
	Duration       string    `json:"duration" gorm:"type:varchar(100) not null; default '0'; comment:'时长' "`
	StartTime      time.Time `json:"start_time" time_format:"2006-01-02 15:04:05" gorm:"comment:'开课时间' " `
	CourseType     int       `json:"course_type" gorm:"default '1';comment:'课程类型 1一对多 2一对一'" `
	ChargeType     int       `json:"charge_type" gorm:"default '1';comment:'收费方式 1按课时 2按月收费'"`
	CourseMethodId int       `json:"course_method_id" gorm:"default '1';comment:'收费方式 1按课时 2按月收费'" `
	CourseRemark   string    `json:"course_remark" gorm:"type:text; comment:'备注' " `
	ScheduleId     int       `json:"schedule_id" gorm:"default '0'; comment:'上课安排'"`

	erm.Model
}
