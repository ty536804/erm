package School

import (
	erm "erm/Database"
)

// @Summer school name
type School struct {
	Id         int    `json:"id" gorm:"primary_key"`
	SchoolName string `json:"school_name"  gorm:"type:varchar(100);not null; default ''; comment:'学校名称' " `
	CourseId   int    `json:"course_id" gorm:"comment:'课程' " `
	SClassId   int    `json:"s_class_id" gorm:"comment:'班级ID' " `

	erm.Model
}
